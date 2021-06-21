package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	res := math.MaxInt64

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		depth++

		if root.Left == nil && root.Right == nil {
			res = min(res, depth)
		}

		dfs(root.Left, depth)
		dfs(root.Right, depth)
	}

	dfs(root, 0)

	if res < math.MaxInt64 {
		return res
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
