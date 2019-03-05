package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func TreeDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	var left int
	var right int

	if root.Left != nil {
		left = TreeDepth(root.Left)
	}
	if root.Right != nil {
		right = TreeDepth(root.Right)
	}

	var ret int
	if left >= right {
		ret = left + 1
	} else {
		ret = right + 1
	}

	return ret
}

func main() {
	root := &BinaryTreeNode{Value: 1}
	left := &BinaryTreeNode{Value: 2}
	right := &BinaryTreeNode{Value: 3}

	root.Left = left
	root.Right = right

	left1 := &BinaryTreeNode{Value: 4}
	right1 := &BinaryTreeNode{Value: 5}

	left.Left = left1
	left1.Right = right1

	depth := TreeDepth(root)
	fmt.Printf("The binary tree's depth: %d\n", depth)
}
