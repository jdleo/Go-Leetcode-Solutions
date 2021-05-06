package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := []int{}

	// go while queue is non empty
	for len(queue) > 0 {
		// get current size of current level
		size := len(queue)
		// rightmost node val
		rightmost := -1
		// go through current level
		for i := 0; i < size; i++ {
			// popleft from queue
			cur := queue[0]
			queue = queue[1:]
			// update rightmost
			rightmost = cur.Val
			// add children to queue
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// add rightmost to res
		res = append(res, rightmost)
	}

	return res
}
