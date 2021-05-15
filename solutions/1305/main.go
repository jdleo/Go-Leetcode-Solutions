package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	// get inorders of root1, root2
	slice1 := inorder(root1)
	slice2 := inorder(root2)
	// make result slice of the lengths combined
	res := make([]int, len(slice1)+len(slice2))
	// go through both slices
	i, j, ptr := 0, 0, 0
	for i < len(slice1) || j < len(slice2) {
		// take the smaller one
		if i == len(slice1) {
			res[ptr] = slice2[j]
			j++
		} else if j == len(slice2) {
			res[ptr] = slice1[i]
			i++
		} else if slice1[i] < slice2[j] {
			res[ptr] = slice1[i]
			i++
		} else {
			res[ptr] = slice2[j]
			j++
		}

		ptr++
	}
	return res
}

// helper method to get inorder traversal of a tree
func inorder(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		if root != nil {
			// keep going left, adding to stack
			stack = append(stack, root)
			root = root.Left
		} else {
			// pop from stack
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 'log' this value
			res = append(res, root.Val)
			// go right
			root = root.Right
		}
	}

	return res
}
