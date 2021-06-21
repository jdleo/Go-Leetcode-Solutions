package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	backtrack(&res, &[]int{}, root, targetSum)
	return res
}

func backtrack(res *[][]int, path *[]int, root *TreeNode, targetSum int) {
	if root == nil {
		return
	}

	targetSum -= root.Val
	*path = append(*path, root.Val)

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		temp := make([]int, len(*path))
		copy(temp, *path)
		*res = append(*res, temp[:])
	}

	backtrack(res, path, root.Left, targetSum)
	backtrack(res, path, root.Right, targetSum)
	*path = (*path)[:len(*path)-1]
}
