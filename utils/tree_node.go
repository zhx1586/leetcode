package util

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree(input string) *TreeNode {
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	if len(input) == 0 {
		return nil
	}
	vals := strings.Split(input, ",")

	dummyRoot := &TreeNode{}
	queue := []*TreeNode{dummyRoot}
	for i, val := range vals {
		node := queue[0]
		if i%2 == 0 {
			queue = queue[1:]
		}
		if val == "null" {
			continue
		}

		n, _ := strconv.Atoi(val)
		child := &TreeNode{Val: n}
		if i%2 != 0 {
			node.Left = child
		} else {
			node.Right = child
		}
		queue = append(queue, child)
	}
	return dummyRoot.Right
}
