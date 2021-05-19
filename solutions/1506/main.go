package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func findRoot(tree []*Node) *Node {
	// map to track if nodes have a parent
	hasParent := map[*Node]bool{}

	// go through tree list
	for _, node := range tree {
		// mark all children as having a parent
		for _, child := range node.Children {
			hasParent[child] = true
		}
	}

	// go back through tree list
	for _, node := range tree {
		// check if has no parent (this is root)
		if !hasParent[node] {
			return node
		}
	}

	return nil
}
