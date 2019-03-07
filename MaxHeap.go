package main

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	Element []int
}

func NewMaxHeap() *MaxHeap {
	h := &MaxHeap{Element: []int{math.MaxInt64}}
	return h
}

func (h *MaxHeap) Insert(v int) {
	h.Element = append(h.Element, v)

	i := len(h.Element) - 1
	for ; h.Element[i/2] < v; i /= 2 {
		h.Element[i] = h.Element[i/2]
	}

	h.Element[i] = v
}

func (h *MaxHeap) DeleteMax() (int, error) {
	if len(h.Element) <= 1 {
		return 0, fmt.Errorf("MaxHeap is empty")
	}

	maxElement := h.Element[1]
	lastElement := h.Element[len(h.Element)-1]

	var i, child int
	for i = 1; i*2 < len(h.Element); i = child {
		child = i * 2
		if child < len(h.Element)-1 && h.Element[child+1] > h.Element[child] {
			child++
		}

		if lastElement < h.Element[child] {
			h.Element[i] = h.Element[child]
		} else {
			break
		}
	}

	h.Element[i] = lastElement
	h.Element = h.Element[:len(h.Element)-1]
	return maxElement, nil
}

func main() {
	h := NewMaxHeap()
	h.Insert(8)
	h.Insert(1)
	h.Insert(4)
	h.Insert(5)
	h.Insert(2)

	max, err := h.DeleteMax()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(max)

	max, err = h.DeleteMax()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(max)
}
