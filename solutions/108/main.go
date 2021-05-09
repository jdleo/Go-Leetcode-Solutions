package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// pass to recursive builder func
	return builder(nums, 0, len(nums)-1)
}

func builder(nums []int, left, right int) *TreeNode {
	// base case
	if left > right {
		return nil
	}

	// get mid point
	mid := (left + right) >> 1

	// current root node is num at this mid point
	root := &TreeNode{nums[mid], nil, nil}

	// build left, right
	root.Left = builder(nums, left, mid-1)
	root.Right = builder(nums, mid+1, right)

	return root
}
