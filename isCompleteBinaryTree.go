package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type Queue struct {
	root []*BinaryTreeNode
}

func (q *Queue) Push(node *BinaryTreeNode) {
	q.root = append(q.root, node)
}

func (q *Queue) Pop() *BinaryTreeNode {
	res := q.root[0]
	q.root = q.root[1:]
	return res
}

func (q *Queue) Empty() bool {
	return len(q.root) == 0
}

func IsBinaryTree(root *BinaryTreeNode) bool {
	if root == nil {
		return false
	}

	queue := Queue{root:make([]*BinaryTreeNode, 0)}
	queue.Push(root)

	for node := queue.Pop(); node != nil; node = queue.Pop() {
		queue.Push(node.Left)
		queue.Push(node.Right)
	}

	for !queue.Empty() {
		item := queue.Pop()

		if item != nil {
			return false
		}
	}

	return true
}

func main() {
	root := &BinaryTreeNode{Value: 1}
	//left := &BinaryTreeNode{Value: 2}
	right := &BinaryTreeNode{Value: 3}

	//root.Left = left
	root.Right = right
	
	if IsBinaryTree(root) {
		fmt.Println("Is a complete binary tree!")
		return
	}
	fmt.Println("Not a complete binary tree!")
}
