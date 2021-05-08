package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	res := 0
	maxDepth := 0

	// run helper
	helper(&res, &maxDepth, root, 0)

	return res
}

func helper(res *int, maxDepth *int, root *TreeNode, depth int) {
	// base case
	if root == nil {
		return
	}

	// check if current depth is at the max depth
	if depth == *maxDepth {
		// add to res
		*res += root.Val
	} else if depth > *maxDepth {
		// we need to update new max depth and res
		*maxDepth = depth
		*res = root.Val
	}

	// go left and right
	helper(res, maxDepth, root.Left, depth+1)
	helper(res, maxDepth, root.Right, depth+1)
}
