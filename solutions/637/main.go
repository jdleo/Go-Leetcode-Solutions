package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	// queue for bfs
	queue := []*TreeNode{root}

	// go until empty queue
	for len(queue) != 0 {
		// current level's size, and sum
		size, sum := len(queue), 0

		// go thru current level
		for i := 0; i < size; i++ {
			// popleft, add to sum
			cur := queue[0]
			queue = queue[1:]
			sum += cur.Val

			// add children to queue
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		// push average
		res = append(res, float64(sum)/float64(size))
	}

	return res
}
