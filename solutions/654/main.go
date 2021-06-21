package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// call recursive builder
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	// base case : left and right crossover (no nums to search)
	if left > right {
		return nil
	}

	// find max index in nums
	maxIdx := left
	for i := maxIdx; i <= right; i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	// make new root with num at max index
	root := &TreeNode{nums[maxIdx], nil, nil}

	// set left and right
	root.Left = build(nums, left, maxIdx-1)
	root.Right = build(nums, maxIdx+1, right)

	return root
}
