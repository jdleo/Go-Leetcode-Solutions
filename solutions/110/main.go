package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res := true
	maxDepth(root, &res)
	return res
}

func maxDepth(root *TreeNode, res *bool) int {
	if !*res {
		return 0
	} else if root == nil {
		return 0
	}

	left, right := maxDepth(root.Left, res), maxDepth(root.Right, res)

	if abs(left-right) > 1 {
		*res = false
	}

	return 1 + max(left, right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
