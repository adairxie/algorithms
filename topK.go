package main

import (
	"fmt"
	"math/rand"
)

// 随机生成一个index, index的左边是比它的值，右边是比它小的值
func partition(nums []int, start, end int) int {
	if start >= end {
		return start
	}

	index := start + rand.Intn(end-start+1)
	nums[index], nums[end] = nums[end], nums[index]

	small := start - 1
	for index = start; index < end; index++ {
		if nums[index] > nums[end] {
			small++
			if small != index {
				nums[small], nums[index] = nums[index], nums[small]
			}
		}
	}
	small++
	nums[small], nums[end] = nums[end], nums[small]

	return small
}

func topK(nums []int, k int) []int {
	if nums == nil || len(nums) <= 0 {
		return nil
	}

	start := 0
	end := len(nums) - 1
	index := partition(nums, start, end)

	for index != k-1 {
		if index > start {
			index = partition(nums, start, index-1)
		}

		if index < end {
			index = partition(nums, index+1, end)
		}
	}

	return nums[0:k]
}

func main() {
	array := []int{8, 1, 2, 4, 4}

	res := topK(array, 3)
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
