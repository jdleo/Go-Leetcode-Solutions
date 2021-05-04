package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}

	// get left and right subtree depths
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// check max
	if left > right {
		// count current depth + left
		return 1 + left
	} else {
		// count current depth + right
		return 1 + right
	}
}
