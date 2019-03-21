package main

import "fmt"

func mergeSort(data []int, start, end int) {
	if data == nil || start >= end || start < 0 || end < 0 {
		return
	}

	mid := (start + end) / 2
	mergeSort(data, start, mid)
	mergeSort(data, mid+1, end)
	merge(data, start, mid, end)
}

func merge(data []int, start, mid, end int) {
	if start == end {
		return
	}

	tmp := []int{}
	i := start
	j := mid + 1

	for i <= mid && j <= end {
		if data[i] < data[j] {
			tmp = append(tmp, data[i])
			i++
		} else {
			tmp = append(tmp, data[j])
			j++
		}
	}

	for i <= mid {
		tmp = append(tmp, data[i])
		i++
	}

	for j <= end {
		tmp = append(tmp, data[j])
		j++
	}

	for index, val := range tmp {
		data[start + index] = val
	}
}

func main() {
	arr := []int{2, 4, 6, 0, 9, 8}
	mergeSort(arr, 0, len(arr)-1)
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
