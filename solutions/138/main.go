package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// map to hold nodes in original with their copies
	hm := map[*Node]*Node{}

	// first, go through and make a copy of all nodes vals
	cur := head
	for cur != nil {
		hm[cur] = &Node{cur.Val, nil, nil}
		cur = cur.Next
	}

	// next, go back through list
	cur = head
	for cur != nil {
		// set next and random pointers based on our map
		hm[cur].Next = hm[cur.Next]
		hm[cur].Random = hm[cur.Random]
		cur = cur.Next
	}

	// return the copied head node
	return hm[head]
}
