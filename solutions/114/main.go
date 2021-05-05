package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	cur := root

	for cur != nil {
		if cur.Left != nil {
			// find current node's previous node that links to current node's right
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			// use current node's left to replace its right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}
