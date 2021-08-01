package pathsumiii

import . "github.com/zhx1586/leetcode/utils"

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	ret := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		pathSumCore(curr, 0, targetSum, &ret)
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return ret
}

func pathSumCore(curr *TreeNode, sum int, targetSum int, ret *int) {
	if curr == nil {
		return
	}
	if sum+curr.Val == targetSum {
		*ret = *ret + 1
	}
	pathSumCore(curr.Left, sum+curr.Val, targetSum, ret)
	pathSumCore(curr.Right, sum+curr.Val, targetSum, ret)
}
