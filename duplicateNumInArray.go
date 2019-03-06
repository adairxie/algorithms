//在一个长度为n的数组中存放着0~n-1的数字
//找出数组中第一个重复的数字

package main

import "fmt"

func duplicate(data []int) (int, bool) {
	for i := 0; i < len(data); i++ {
		for data[i] != i {
			if data[i] == data[data[i]] {
				return data[i], true
			}

			data[i], data[data[i]] = data[data[i]], data[i]
		}
	}

	return -1, false
}

func main() {
	array := []int{1, 2, 3, 4, 3, 5}
	du, ok := duplicate(array)
	if ok {
		fmt.Printf("duplicate number: %d\n", du)
		return
	}

	fmt.Println("no duplicate number!")
}
