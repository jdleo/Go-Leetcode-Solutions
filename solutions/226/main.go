package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// swap
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
