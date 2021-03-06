package main

import (
	"fmt"
    "sync"
)

func PrintStringSlice(s []string) {
	for _, v := range s {
		fmt.Printf("%s\t", v)
	}
	fmt.Println()
}

func parrell(data []string, chunk, parrell int, processor func([]string)) {
	ch := make(chan struct{}, parrell)
    var wg sync.WaitGroup

	data_len := len(data)
	for idx := 0; idx < data_len; idx += chunk {
        wg.Add(1)
		if idx+chunk >= data_len {
			msg := data[idx:data_len]
			ch <- struct{}{}
			go func(msg []string) {
                defer wg.Done()
				processor(msg)
				<-ch
			}(msg)
		} else {
			msg := data[idx : idx+chunk]
			ch <- struct{}{}
			go func(msg []string) {
                defer wg.Done()
				processor(msg)
				<-ch
			}(msg)
		}
	}

    wg.Wait()
}

func main() {
	str := []string{"1", "2", "3", "4", "5", "6", "7"}
	parrell(str, 3, 2, PrintStringSlice)
	fmt.Println("end")
}
