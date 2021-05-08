package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	// start max depth
	maxDepth(root, &res)
	return res
}

// helper function for calculating max depth and updating result
func maxDepth(root *TreeNode, res *int) int {
	// base case
	if root == nil {
		return 0
	}

	// get left and right max depths
	left := maxDepth(root.Left, res)
	right := maxDepth(root.Right, res)

	// update res
	if left+right > *res {
		*res = left + right
	}

	// return 1 + max depth (counting this level)
	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
}
