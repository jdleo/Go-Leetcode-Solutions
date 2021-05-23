package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
	root  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{[]*TreeNode{}, root}
}

func (this *BSTIterator) Next() int {
	// go until root is nil (keep going left)
	for this.root != nil {
		// add to stack and go left
		this.stack = append(this.stack, this.root)
		this.root = this.root.Left
	}

	// pop from stack, return inorder, and go right
	this.root = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	res := this.root.Val
	this.root = this.root.Right
	return res
}

func (this *BSTIterator) HasNext() bool {
	// check if stack or root is not nil
	return len(this.stack) > 0 || this.root != nil
}
