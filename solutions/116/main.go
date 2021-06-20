package main

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// make queue for bfs
	queue := []*Node{root}

	for len(queue) != 0 {
		// popleft
		cur := queue[0]
		queue = queue[1:]

		// check if cur has left and right to connect
		if cur.Left != nil && cur.Right != nil {
			// connect left to right
			cur.Left.Next = cur.Right

			// check if this has a next (to connect across trees)
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}

			// add children to queue
			queue = append(queue, cur.Left, cur.Right)
		}
	}

	return root
}
