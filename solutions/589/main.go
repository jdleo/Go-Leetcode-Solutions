package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := []int{}
	// start recursive preorder function
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	// base case
	if root == nil {
		return
	}

	// visit this node before exploring others (PREorder)
	*res = append(*res, root.Val)

	// visit children
	for _, child := range root.Children {
		helper(child, res)
	}
}
