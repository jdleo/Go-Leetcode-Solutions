package main

// Definition for a Node.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Random *Node
}

type NodeCopy struct {
	Val    int
	Left   *NodeCopy
	Right  *NodeCopy
	Random *NodeCopy
}

func copyRandomBinaryTree(root *Node) *NodeCopy {
	if root == nil {
		return nil
	}

	// map to map original nodes to their node copies
	hm := map[*Node]*NodeCopy{}
	// stack for dfs
	stack := []*Node{root}
	// dfs
	for len(stack) != 0 {
		// pop
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// make a copy of this node
		hm[cur] = &NodeCopy{cur.Val, nil, nil, nil}

		// add children
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	// do dfs again
	stack = []*Node{root}
	for len(stack) != 0 {
		// pop
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// get the copies of left, right, random
		hm[cur].Left = hm[cur.Left]
		hm[cur].Right = hm[cur.Right]
		hm[cur].Random = hm[cur.Random]

		// add children
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	// get the root node from hashmap
	return hm[root]
}
