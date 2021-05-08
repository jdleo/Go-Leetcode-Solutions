package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	// start helper and take max path
	a, b := helper(root)
	return max(a, b)
}

// recursive helper
func helper(root *TreeNode) (a int, b int) {
	// base case
	if root == nil {
		return 0, 0
	}
	// calculate left and right subtree paths
	l_skip, l_rob := helper(root.Left)
	r_skip, r_rob := helper(root.Right)

	// this is the result if we are to rob right here
	rob := root.Val + l_skip + r_skip
	// this is the result if we are to skip here
	skip := max(l_skip, l_rob) + max(r_skip, r_rob)

	return skip, rob
}

// helper function for max
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
