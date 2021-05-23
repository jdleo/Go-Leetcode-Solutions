package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue, level, maxSum, maxLevel := []*TreeNode{root}, 1, math.MinInt64, -1

	// go until empty queue
	for len(queue) != 0 {
		// get size of queue
		size := len(queue)
		// current level sum
		sum := 0

		// go through level
		for i := 0; i < size; i++ {
			// popleft, add sum
			cur := queue[0]
			queue = queue[1:]
			sum += cur.Val

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		// set new max sum
		if sum > maxSum {
			maxSum = sum
			maxLevel = level
		}

		level++
	}

	return maxLevel
}
