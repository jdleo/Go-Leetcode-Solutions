package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	// queue for bfs
	queue := []*TreeNode{root}
	// result node
	res := root.Val

	// go until empty queue
	for len(queue) != 0 {
		// get size of level
		size := len(queue)
		// go through level
		for i := 0; i < size; i++ {
			// popleft
			cur := queue[0]
			queue = queue[1:]
			// check if leftmost of this level
			if i == 0 {
				res = cur.Val
			}

			// add left, right
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return res
}
