package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// base case : both roots are null
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 != nil && root2 != nil {
		// build left and right subtrees as normal
		left := mergeTrees(root1.Left, root2.Left)
		right := mergeTrees(root1.Right, root2.Right)

		// return merged node
		return &TreeNode{root1.Val + root2.Val, left, right}
	} else if root1 != nil {
		// build left and right subtrees (but root2 is nil)
		left := mergeTrees(root1.Left, nil)
		right := mergeTrees(root1.Right, nil)

		// return merged node (just root1's val, since root2 is nil)
		return &TreeNode{root1.Val, left, right}
	} else {
		// build left and right subtrees (but root1 is nil)
		left := mergeTrees(nil, root2.Left)
		right := mergeTrees(nil, root2.Right)

		// return merged node (just root1's val, since root2 is nil)
		return &TreeNode{root2.Val, left, right}
	}
}
