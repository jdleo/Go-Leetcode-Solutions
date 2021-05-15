package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	// base case
	if root == nil {
		return 0
	}

	// check if this root is in range
	if low <= root.Val && root.Val <= high {
		// we can add this, and also search left, right
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	} else if root.Val < low {
		// we only want to search right
		return rangeSumBST(root.Right, low, high)
	} else {
		// we only want to search left
		return rangeSumBST(root.Left, low, high)
	}
}
