package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// delete in left and subtrees first
	// because we wanna see if after deletion
	// the parent is a leaf
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	// check if this is a leaf and needs to be deleted
	if root.Val == target && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}
