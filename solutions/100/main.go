package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// check if one is nil, or values arent equal
	if p == nil || q == nil {
		return p == q
	} else if p.Val != q.Val {
		return false
	}

	// get left and right
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
