package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// stack for inorder traversal
	stack := []*TreeNode{}

	// go while root is not nil, or theres something in stack
	for root != nil || len(stack) > 0 {
		// check if root is not nil
		if root != nil {
			// go left
			stack = append(stack, root)
			root = root.Left
		} else {
			// pop from stack
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// decrement k, and check if this is kth smallest
			k--
			if k == 0 {
				return root.Val
			}
			// go right
			root = root.Right
		}
	}

	// never found, idk
	return -1
}
