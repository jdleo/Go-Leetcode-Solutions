package main

// Definition for Node.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	// find root of tree
	root := getRoot(p)
	// parent mapping and queue for bfs
	parent := map[*Node]*Node{}
	queue := []*Node{root}

	// go while queue isnt empty
	for len(queue) > 0 {
		// offer from queue
		cur := queue[0]
		queue = queue[1:]

		if cur != nil {
			// map children to parent
			parent[cur.Left] = cur
			parent[cur.Right] = cur
			// add children to queue
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}

	// seen set
	seen := map[*Node]bool{}

	// go through p's path
	for p != nil {
		seen[p] = true
		p = parent[p]
	}

	// go until we find q in seen (LCA)
	for !seen[q] {
		q = parent[q]
	}

	return q
}

// helper method to find root
func getRoot(node *Node) *Node {
	for node.Parent != nil {
		node = node.Parent
	}
	return node
}
