package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLonelyNodes(root *TreeNode) []int {
	res := []int{}
	// start recursive dfs helper
	dfs(&res, root)
	return res
}

func dfs(res *[]int, root *TreeNode) {
	// base case
	if root == nil {
		return
	}

	// check lonely node
	if root.Left == nil && root.Right != nil {
		*res = append(*res, root.Right.Val)
	} else if root.Left != nil && root.Right == nil {
		*res = append(*res, root.Left.Val)
	}

	// go left and right
	dfs(res, root.Left)
	dfs(res, root.Right)
}
