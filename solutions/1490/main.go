package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	// base case
	if root == nil {
		return nil
	}

	// clone this node
	clone := &Node{root.Val, []*Node{}}

	// go through children in node
	for _, child := range root.Children {
		// clone individual child
		clone.Children = append(clone.Children, cloneTree(child))
	}

	return clone
}
