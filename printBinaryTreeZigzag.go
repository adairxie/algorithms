package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type stack struct {
	root []*BinaryTreeNode
}

func (s *stack) Push(node *BinaryTreeNode) {
	if s.root == nil {
		s.root = make([]*BinaryTreeNode, 0)
	}
	s.root = append(s.root, node)
}

func (s *stack) Pop() *BinaryTreeNode {
	res := s.root[len(s.root)-1]
	s.root = s.root[:(len(s.root) - 1)]
	return res
}

func (s *stack) Empty() bool {
	return len(s.root) == 0
}

func printBinaryTreeZigZag(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	s := [2]stack{}

	current := 0
	next := 1
	s[current].Push(root)

	for !s[0].Empty() || !s[1].Empty() {
		node := s[current].Pop()
		fmt.Printf("%d ", node.Value)

		if current == 0 {
			if node.Left != nil {
				s[next].Push(node.Left)
			}
			if node.Right != nil {
				s[next].Push(node.Right)
			}
		} else {
			if node.Right != nil {
				s[next].Push(node.Right)
			}
			if node.Left != nil {
				s[next].Push(node.Left)
			}
		}

		if s[current].Empty() {
			fmt.Println()
			current = 1 - current
			next = 1 - next
		}

	}
}

func main() {
	root := &BinaryTreeNode{Value: 1}
	left := &BinaryTreeNode{Value: 2}
	right := &BinaryTreeNode{Value: 3}

	root.Left = left
	root.Right = right

	left1 := &BinaryTreeNode{Value: 4}
	left.Left = left1
	right1 := &BinaryTreeNode{Value: 5}
	left.Right = right1

	left2 := &BinaryTreeNode{Value: 6}
	right.Left = left2
	right2 := &BinaryTreeNode{Value: 7}
	right.Right = right2

	printBinaryTreeZigZag(root)
}
