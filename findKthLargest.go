//Find the kth largest element in an unsorted array.
//Note that it is the kth largest element in the sorted order, not the kth distinct element.
// Example 1:
//   Input: [3, 2, 1, 5, 6, 4] and k = 2
//   Output: 5

// Example 2:
//  Input: [3, 2, 3, 1, 2, 4,5 ,5, 6] and k = 4
//  Output: 4

package main

import (
	"fmt"
	"math/rand"
)

func Print(nums []int) {
	for _, e := range nums {
		fmt.Printf("%d ", e)
	}

	fmt.Println()
}

func findKthLargest(nums []int, k int) int {
	dst_index := len(nums) - k

	start := 0
	end := len(nums) - 1

	index := partition(nums, start, end)

	for index != dst_index {
		if index < dst_index {
			index = partition(nums, index+1, end)
		}

		if index > dst_index {
			index = partition(nums, start, index-1)
		}
	}

	return nums[index]
}

func partition(nums []int, start, end int) int {
	index := start + rand.Intn(end - start + 1)

	small := start - 1
	nums[index], nums[end] = nums[end], nums[index]

	for i := start; i <= end; i++ {
		if nums[i] < nums[end] {
			small = small + 1
			if small != i {
				nums[small], nums[i] = nums[i], nums[small]
			}

		}

	}
	small = small + 1
	nums[small], nums[end] = nums[end], nums[small]

	return small
}

func main() {
	input1 := []int{2, 1}

	KthElem := findKthLargest(input1, 2)
	fmt.Println(KthElem)
}
