package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	res := [][]int{}
	helper(root, &res)
	return res
}

// recursive helper function
func helper(root *TreeNode, res *[][]int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, res)
	right := helper(root.Right, res)
	level := max(left, right) + 1

	if level > len(*res) {
		*res = append(*res, []int{})
	}

	(*res)[level-1] = append((*res)[level-1], root.Val)

	return level
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
