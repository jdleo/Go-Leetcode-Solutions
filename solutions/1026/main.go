package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	// start recursion
	return helper(root, root.Val, root.Val)
}

// recursive helper method to keep track of min/max and max diff
func helper(root *TreeNode, mn int, mx int) int {
	if root == nil {
		// greatest difference in this path
		return mx - mn
	}

	// update min/max
	if root.Val < mn {
		mn = root.Val
	}
	if root.Val > mx {
		mx = root.Val
	}

	// calculate left and right
	left := helper(root.Left, mn, mx)
	right := helper(root.Right, mn, mx)

	// return bigger
	if left > right {
		return left
	} else {
		return right
	}
}
