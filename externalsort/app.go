package main

import (
	"bufio"
	"container/heap"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	BlockBufferSize = 0
	CacheBufferSize = 0
	LineBufferSize  = 8

	ReadWriteBufferPercentage = 5
	BlockPercentage           = 90

	SLEEP_TIME_MS = 100
	CACHE_PREFIX  = "cache/block_"
)

var memoryBufferSize int64 = 0

// Times
var iowaitRead = time.Duration(0)
var iowaitWrite = time.Duration(0)
var timeSlept = time.Duration(0)
var timeSorting = time.Duration(0)
var timeMerging = time.Duration(0)

func readLines(srcFile *os.File, blockCh chan<- *[]string) {
	defer srcFile.Close()
	defer close(blockCh)

	var bufferSize int64 = (Options.BufferSize / 100) * ReadWriteBufferPercentage
	var blockSize int64 = (Options.BufferSize / 100) * BlockPercentage

	scanner := bufio.NewScanner(srcFile)
	scanner.Buffer(make([]byte, bufferSize), 0)

	var block *[]string
	var blockBufferSize int64

	// stop reading file on SIGINT
	for !sigINT {
		// sleep if memory buffer is full
		for atomic.LoadInt64(&memoryBufferSize) >= blockSize {
			start := time.Now()
			Sleepms(SLEEP_TIME_MS)
			timeSlept += time.Since(start)
		}

		start := time.Now()
		scan := scanner.Scan()
		iowaitRead += time.Since(start)

		// EOF (or error)
		if !scan {
			break
		}

		// create new block
		if block == nil {
			block = new([]string)
			blockBufferSize = 0
		}

		*block = append(*block, scanner.Text())
		blockBufferSize += int64(len(scanner.Bytes()))
		atomic.AddInt64(&memoryBufferSize, int64(len(scanner.Bytes())))

		if blockBufferSize >= blockSize {
			blockCh <- block
			block = nil
		}
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		Stderrln("Error: ", err)
	}

	if len(*block) > 0 {
		blockCh <- block
		block = nil
	}
}

func sortBlocks(blockCh <-chan *[]string, saveCh chan<- *[]string) {
	defer close(saveCh)

	for block := range blockCh {
		start := time.Now()
		if Options.Quicksort {
			sort.Sort(Alphabetical(*block))
		} else {
			Heapsort(Alphabetical(*block))
		}
		timeSorting += time.Since(start)
		saveCh <- block
	}
}

func saveBlocks(tempFiles *[]string, saveCh <-chan *[]string, wg *sync.WaitGroup) {
	defer wg.Done()

	var line string
	bufferSize := (int(Options.BufferSize) / 100) * ReadWriteBufferPercentage

	idx := 0
	for block := range saveCh {
		// create temp file
		tempFile, err := os.Create(CACHE_PREFIX + strconv.Itoa(idx))
		if err != nil {
			Stderrln("Error: ", err)
		}
		writer := bufio.NewWriterSize(tempFile, bufferSize)

		// keep track of temp file
		*tempFiles = append(*tempFiles, tempFile.Name())

		for len(*block) > 0 {
			line = (*block)[0]
			*block = (*block)[1:]

			// write line to temp file
			start := time.Now()
			writer.WriteString(line)
			writer.WriteString("\n")
			iowaitWrite += time.Since(start)

			atomic.AddInt64(&memoryBufferSize, int64(-1*len(line)))
		}

		// empty the buffer
		start := time.Now()
		writer.Flush()
		iowaitWrite += time.Since(start)

		tempFile.Close()

		idx++
	}
}

func readBlock(tempFile *os.File, bufferSize int, lineCh chan<- string) {
	defer tempFile.Close()
	defer close(lineCh)

	scanner := bufio.NewScanner(tempFile)
	scanner.Buffer(make([]byte, bufferSize), 0)

	// stop reading file on SIGINT
	for !sigINT {
		start := time.Now()
		scan := scanner.Scan()
		iowaitWrite += time.Since(start)

		if !scan {
			break
		}

		// append line
		lineCh <- scanner.Text()
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		Stderrln("Error: ", err)
	}
}

func mergeLines(lineChannels map[int]chan string, mergeCh chan<- string) {
	defer close(mergeCh)

	var lHeap LineHeap
	heap.Init(&lHeap)

	// - one element from each channel to the heap
	for idx, lineCh := range lineChannels {
		line := Line{value: <-lineCh, idx: idx} // bug
		start := time.Now()
		heap.Push(&lHeap, line)
		timeMerging += time.Since(start)
	}

	for lHeap.Len() > 0 {
		firstLine := heap.Pop(&lHeap).(Line)
		mergeCh <- firstLine.value

		line, ok := <-lineChannels[firstLine.idx]
		if ok {
			newLine := Line{value: line, idx: firstLine.idx}
			start := time.Now()
			heap.Push(&lHeap, newLine)
			timeMerging += time.Since(start)
		}
	}
}

func writeFile(destFile *os.File, bufferSize int, mergeCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer destFile.Close()

	writer := bufio.NewWriterSize(destFile, bufferSize)

	for line := range mergeCh {
		// Write line to dest file
		start := time.Now()
		writer.WriteString(line)
		writer.WriteString("\n")
		iowaitWrite += time.Since(start)
	}

	// empty the buffer
	start := time.Now()
	writer.Flush()
	iowaitWrite += time.Since(start)
}

// App construct an app for external sort
func App() error {
	var wg sync.WaitGroup
	var tempFiles []string
	blockCh := make(chan *[]string, BlockBufferSize)
	saveCh := make(chan *[]string, CacheBufferSize)
	lineChannels := make(map[int]chan string)
	mergeCh := make(chan string, LineBufferSize)

	// open source file
	srcFile, err := os.Open(Options.SourceFile)
	if Iserror(err) {
		return err
	}

	// read lines into blocks from the source file
	go readLines(srcFile, blockCh)

	// sort the lines in the blocks
	go sortBlocks(blockCh, saveCh)

	// save blocks to disk into temporary files
	wg.Add(1)
	go saveBlocks(&tempFiles, saveCh, &wg)
	Stderrln("1st pass: Splitting the file into smaller sorted temporary files")
	wg.Wait()

	// Read lines from sorted blocks, put them in channels
	bufferSize := int(Options.BufferSize) / (len(tempFiles) + 1)
	for idx, tempFilePath := range tempFiles {
		lineChannels[idx] = make(chan string, LineBufferSize)

		// open temp file
		tempFile, err := os.Open(tempFilePath)
		if Iserror(err) {
			return err
		}
		go readBlock(tempFile, bufferSize, lineChannels[idx])
	}

	// merge sort lines
	go mergeLines(lineChannels, mergeCh)

	// open destination file
	destFile, err := os.Open(Options.DestFile)
	if Iserror(err) {
		return err
	}
	// Write output file
	wg.Add(1)
	go writeFile(destFile, bufferSize, mergeCh, &wg)
	Stderrln("2nd pass: merging temporary files into destFile")
	wg.Done()

	// cleanup
	for _, file := range tempFiles {
		if !Options.KeepTemps {
			os.Remove(file)
		}
	}

	// print timings
	if Verbose {
		Stderrln("\n(not fully accurate, just for reference)")
		Stderrln("Read iowait: ", iowaitRead)
		Stderrln("Write iowait: ", iowaitWrite)
		Stderrln("Sleep time: ", timeSlept)
		Stderrln("Sorting time: ", timeSorting)
		Stderrln("Merging time: ", timeMerging)
	}

	return nil
}
