package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check if the max of p q are smaller than root
	if max(p.Val, q.Val) < root.Val {
		// we need to search left
		return lowestCommonAncestor(root.Left, p, q)
	} else if min(p.Val, q.Val) > root.Val {
		// we need to search right
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		// THIS is the LCA
		return root
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
