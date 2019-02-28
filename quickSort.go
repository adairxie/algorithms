package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(data []int, start, end int) {
	if data == nil || start < 0 || end < 0 || start > end {
		return
	}

	index := partition(data, start, end)

	if index > start {
		QuickSort(data, start, index-1)
	}

	if index < end {
		QuickSort(data, index+1, end)
	}
}

func partition(data []int, start, end int) int {
	index := start + rand.Intn(end-start+1)
	data[index], data[end] = data[end], data[index]

	small := start - 1

	for i := start; i <= end; i++ {
		if data[i] < data[end] {
			small = small + 1
			if small != i {
				data[small], data[i] = data[i], data[small]
			}
		}
	}
	small = small + 1
	data[end], data[small] = data[small], data[end]

	return small
}

func main() {
	test := []int{1, 3, 5, 2, 6, 4}

	QuickSort(test, 0, len(test)-1)
	for _, elm := range test {
		fmt.Printf("%d ", elm)
	}
	fmt.Println()
}
