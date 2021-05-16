package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	seen map[int]bool // map to hold numbers we've seen
}

func Constructor(root *TreeNode) FindElements {
	findElements := FindElements{map[int]bool{}}
	// recursively build numbers we've seen
	findElements.Dfs(root, 0)
	return findElements
}

// helper method to recursively use dfs to find nodes
func (this *FindElements) Dfs(root *TreeNode, val int) {
	// base case
	if root == nil {
		return
	}

	// add to seen
	this.seen[val] = true

	// go left, right
	this.Dfs(root.Left, 2*val+1)
	this.Dfs(root.Right, 2*val+2)
}

func (this *FindElements) Find(target int) bool {
	// simply return if we've seen this
	return this.seen[target]
}
