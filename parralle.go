package main

import (
	"fmt"
	"time"
)

func PrintStringSlice(s []string) {
	for _, v := range s {
		fmt.Printf("%s\t", v)
	}
	fmt.Println()
}

func parrell(data []string, chunk, parrell int, processor func([]string)) {
	ch := make(chan struct{}, parrell)

	data_len := len(data)
	for idx := 0; idx < data_len; idx += chunk {
		if idx+chunk >= data_len {
			msg := data[idx:data_len]
			ch <- struct{}{}
			go func(msg []string) {
				processor(msg)
				<-ch
			}(msg)
		} else {
			msg := data[idx : idx+chunk]
			ch <- struct{}{}
			go func(msg []string) {
				processor(msg)
				<-ch
			}(msg)
		}
	}
}

func main() {
	str := []string{"1", "2", "3", "4", "5", "6", "7"}
	parrell(str, 3, 2, PrintStringSlice)
	time.Sleep(time.Second * 1)
	fmt.Println("end")
}
