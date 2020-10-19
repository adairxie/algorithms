package main

import (
	"sort"
	"unicode"
	"unicode/utf8"
)

// Return true if string i goes before j
func compare(str_i, str_j string) bool {
	//return str_i < str_j  // lexicographical comparison, faster but crappier

	for len(str_i) > 0 && len(str_j) > 0 {
		ri, size_i := utf8.DecodeRuneInString(str_i)
		rj, size_j := utf8.DecodeRuneInString(str_j)

		ri_lower := unicode.ToLower(ri)
		rj_lower := unicode.ToLower(rj)

		if ri_lower != rj_lower {
			return ri_lower < rj_lower
		}

		str_i = str_i[size_i:]
		str_j = str_j[size_j:]
	}

	return false
}

type Alphabetical []string

// Sort interface
func (l Alphabetical) Len() int { // Return number of items
	return len(l)
}
func (l Alphabetical) Swap(i, j int) { // Swap indexes between items
	l[i], l[j] = l[j], l[i]
}
func (l Alphabetical) Less(i, j int) bool { // Comparision function
	return compare(l[i], l[j])
}

// Line represent a line in file
type Line struct {
	value string
	idx   int
}

// LineHeap a heap with line
type LineHeap []Line

// Sort interface
// Len return heap size
func (l LineHeap) Len() int { // Return number of items
	return len(l)
}

// Swap swap two elements in heap
func (l LineHeap) Swap(i, j int) { // Swap indexes between items
	l[i], l[j] = l[j], l[i]
}
func (l LineHeap) Less(i, j int) bool { // Comparision function
	return compare(l[i].value, l[j].value)
}

// Push implement heap push interface
func (l *LineHeap) Push(x interface{}) { // Append item
	*l = append(*l, x.(Line))
}

// Pop implement heap pop interface
func (l *LineHeap) Pop() interface{} { // Get last item
	hLen := len(*l)
	value := (*l)[hLen-1]
	*l = (*l)[:hLen-1]
	return value
}

// Heapsort implement heap sort algorithm
func Heapsort(data sort.Interface) {
	dataLen := data.Len()

	// Heapify
	for idx := (dataLen - 1) / 2; idx >= 0; idx-- {
		siftDown(data, idx, dataLen)
	}

	// Sorting
	for idx := dataLen - 1; idx >= 0; idx-- {
		data.Swap(0, idx)
		siftDown(data, 0, idx)
	}
}

func siftDown(data sort.Interface, start, end int) {
	root := start

	for {

		child := 2*root + 1

		// Out of index, nothing to do
		if child >= end {
			return
		}

		// Select the greatest child
		if child+1 < end && data.Less(child, child+1) {
			child++
		}

		// Child is smaller than root, nothing to do
		if !data.Less(root, child) {
			return
		}

		// Child is greater than the root, swap them
		data.Swap(root, child)
		root = child
	}
}
