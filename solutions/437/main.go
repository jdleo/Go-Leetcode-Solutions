package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	cache := map[int]int{0: 1}
	helper(&res, cache, root, targetSum, 0)
	return res
}

// recursive helper method
func helper(res *int, cache map[int]int, root *TreeNode, target int, sum int) {
	// base case
	if root == nil {
		return
	}

	// add to running sum
	sum += root.Val
	// check if sum - target prefixsum exists in cache
	if val, ok := cache[sum-target]; ok {
		// add to result (the number of prefix sums is number of valid paths)
		*res += val
	}
	// add this current sum to prefix sums
	cache[sum]++

	// left, right
	helper(res, cache, root.Left, target, sum)
	helper(res, cache, root.Right, target, sum)

	// remove current sum from prefix sums (backtracking way)
	cache[sum]--
}
