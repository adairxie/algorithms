package main

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Element []int
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{Element: []int{math.MinInt64}}
	return h
}

func (h *MinHeap) Insert(v int) {
	h.Element = append(h.Element, v)

	i := len(h.Element) - 1
	for ; h.Element[i/2] > v; i /= 2 {
		h.Element[i] = h.Element[i/2]
	}

	h.Element[i] = v
}

func (h *MinHeap) DeleteMin() (int, error) {
	if len(h.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}

	minElement := h.Element[1]
	lastElement := h.Element[len(h.Element)-1]

	var i, child int
	for i = 1; i*2 < len(h.Element); i = child {
		child = 2 * i
		if child < len(h.Element)-1 && h.Element[child+1] < h.Element[child] {
			child++
		}

		if lastElement > h.Element[child] {
			h.Element[i] = h.Element[child]
		} else {
			break
		}
	}

	h.Element[i] = lastElement

	//删除最后一个元素
	h.Element = h.Element[:len(h.Element)-1]

	return minElement, nil
}

func main() {

	h := NewMinHeap()
	h.Insert(8)
	h.Insert(1)
	h.Insert(4)
	h.Insert(5)
	h.Insert(1)

	min, err := h.DeleteMin()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(min)

	min, err = h.DeleteMin()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(min)

}
