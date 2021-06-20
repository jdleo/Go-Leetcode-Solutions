package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	helper(&res, root)
	return res
}

func helper(res *int, root *TreeNode) int {
	if root == nil {
		return 0
	}

	// calculate left and right, use 0 as bounds (if path is negative)
	left := max(helper(res, root.Left), 0)
	right := max(helper(res, root.Right), 0)

	// set new max
	*res = max(*res, root.Val+left+right)

	// return max path here
	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
