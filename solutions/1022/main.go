package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return helper(root, 0)
}

// recursive helper function
func helper(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}

	// left shift val, and add this current val to it
	val <<= 1
	val |= root.Val

	// check if leaf
	if root.Left == nil && root.Right == nil {
		return val
	}

	// return left + right
	return helper(root.Left, val) + helper(root.Right, val)
}
