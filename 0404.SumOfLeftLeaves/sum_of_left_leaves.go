package sumofleftleaves

import . "github.com/zhx1586/leetcode/utils"

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root != nil {
		sumOfLeftLeavesCore(root, &sum)
	}

	return sum
}

func sumOfLeftLeavesCore(node *TreeNode, sum *int) {
	if node.Left != nil {
		if isLeaf(node.Left) {
			*sum = *sum + node.Left.Val
		} else {
			sumOfLeftLeavesCore(node.Left, sum)
		}
	}
	if node.Right != nil {
		sumOfLeftLeavesCore(node.Right, sum)
	}
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
