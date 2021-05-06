package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	// go while queue isnt empty
	for len(queue) > 0 {
		// get size of current level
		size := len(queue)
		// current level to be filled
		level := []int{}
		// go through current level
		for i := 0; i < size; i++ {
			// popleft from queue
			cur := queue[0]
			queue = queue[1:]
			// add val to level
			level = append(level, cur.Val)
			// add children if non-nil
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// add current level to res
		res = append(res, level)
	}

	return res
}
