package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil // never found
	}

	// check if node is equal to target
	if root.Val == val {
		return root
	}

	// check if less than current (search left)
	if val < root.Val {
		return searchBST(root.Left, val)
	}

	// search right
	return searchBST(root.Right, val)
}
