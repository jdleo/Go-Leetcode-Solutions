package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := map[*TreeNode]*TreeNode{}
	queue := []*TreeNode{root}

	// go while queue is non-empty
	for len(queue) > 0 {
		// popleft
		cur := queue[0]
		queue = queue[1:]

		if cur != nil {
			// make this node the parent for left and right
			parent[cur.Left] = cur
			parent[cur.Right] = cur
			// add left and right to queue
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}

	// seen set
	set := map[*TreeNode]bool{}

	// go through p's path
	for p != nil {
		// add to set, and go thru path
		set[p] = true
		p = parent[p]
	}

	// go through q's path
	for q != nil {
		// check if in set
		if set[q] {
			// first lowest ancestor found
			return q
		}
		q = parent[q]
	}

	return q
}
