package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	// inorder helper
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		// base case
		if root == nil {
			return
		}

		// go left, record val, then right
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}

	// start inorder recursively
	inorder(root)
	return res
}
