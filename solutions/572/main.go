package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	// check if we found potential subroot
	if root.Val == subRoot.Val {
		if isSame(root, subRoot) {
			return true
		}
	}

	// search left and right for root of subroot
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// recursive helper function to determine if two trees are same
func isSame(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}

	// make sure these are the same
	if a.Val != b.Val {
		return false
	}

	return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}
