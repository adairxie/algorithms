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

func printBinaryTree(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	q := &Queue{root: make([]*BinaryTreeNode, 0)}
	q.Push(root)

	nextLevel := 0
	toBePrinted := 1

	for !q.Empty() {
		elem := q.Pop()
		fmt.Printf("%d ", elem.Value)
		if elem.Left != nil {
			q.Push(elem.Left)
			nextLevel++
		}
		if elem.Right != nil {
			q.Push(elem.Right)
			nextLevel++
		}
		toBePrinted--
		if toBePrinted == 0 {
			fmt.Println()
			toBePrinted = nextLevel
			nextLevel = 0
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


	printBinaryTree(root)
}
