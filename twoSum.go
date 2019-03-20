package main

import "fmt"

func twoSum(nums []int, target int) (res []int) {
	intMap := make(map[int]int)

	for idx, num := range nums {
		diff := target - num
		if val, ok := intMap[diff]; ok {
			res = append(res, idx, val)
			return res
		}

		intMap[num] = idx
	}

	return nil
}

func main() {

	arr := []int{2, 7, 11, 5}
	res := twoSum(arr, 9)

	for _, v := range res {
		fmt.Printf("%d\t", v)
	}

	fmt.Println()
}
