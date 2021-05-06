package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// null or start helper
	return root == nil || helper(root.Left, root.Right)
}

// helper method to recursively solve
func helper(left *TreeNode, right *TreeNode) bool {
	// base case
	if left == nil || right == nil {
		// both should be nil
		return left == right
	} else if left.Val != right.Val {
		return false
	}

	// L.L == R.R and L.R = R.L
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
