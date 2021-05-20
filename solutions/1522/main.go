package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {
	res := 0
	// start recursive helper
	helper(root, &res)
	return res
}

// helper function, which gets max depth WHILE updating result
func helper(root *Node, res *int) int {
	// first, and second max depths
	first, second := 0, 0

	// get max depth of all children
	for _, child := range root.Children {
		// get depth
		depth := helper(child, res)
		// set first, second
		if depth > first {
			first, second = depth, first
		} else if depth > second {
			second = depth
		}
	}

	// set new max
	*res = max(*res, first+second)
	// current max depth of THIS node
	return 1 + first
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
