package main

import "fmt"

func mergeSortedArray(arr1, arr2 []int) []int {
	if arr1 == nil {
		return arr2
	}
	if arr2 == nil {
		return arr1
	}

	var res []int
	var i, j int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	if i >= len(arr1) && j < len(arr2) {
		res = append(res, arr2[j:]...)
	}

	if j >= len(arr2) && i < len(arr1) {
		res = append(res, arr1[i:]...)
	}

	return res
}

func main() {
	arr1 := []int{1, 3, 5}
	arr2 := []int{6, 7, 9}

	result := mergeSortedArray(arr1, arr2)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}

	fmt.Println("end")
}
