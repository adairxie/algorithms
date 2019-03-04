package main

import (
	"fmt"
    "errors"
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

func kthNodeBST(root *BinaryTreeNode, k int) (int, error) {
	if root == nil || k <= 0 {
		return 0, errors.New("invalid paramas")
	}

	target := kthNodeBSTCore(root, &k)
    return target.Value, nil
}

func kthNodeBSTCore(root *BinaryTreeNode, k *int) *BinaryTreeNode {
	var target *BinaryTreeNode
	if root.Left != nil {
		target = kthNodeBSTCore(root.Left, k)
	}

	if target == nil {
		if *k == 1 {
			target = root
		}
		(*k)--
	}

	if target == nil && root.Right != nil {
		target = kthNodeBSTCore(root.Right, k)
	}

	return target
}

func main() {
	root := &BinaryTreeNode{Value: 4}
	left := &BinaryTreeNode{Value: 2}
	right := &BinaryTreeNode{Value: 8}

	root.Left = left
	root.Right = right

	left1 := &BinaryTreeNode{Value: 1}
	left.Left = left1
	right1 := &BinaryTreeNode{Value: 3}
	left.Right = right1

	left2 := &BinaryTreeNode{Value: 5}
	right.Left = left2
	right2 := &BinaryTreeNode{Value: 7}
	right.Right = right2

	res, err := kthNodeBST(root, 6)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
