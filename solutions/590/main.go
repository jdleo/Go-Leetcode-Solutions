package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}

	// start recursive helper
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	// explore all children first
	for _, child := range root.Children {
		helper(child, res)
	}

	// add this node val AFTER (postorder)
	*res = append(*res, root.Val)
}
