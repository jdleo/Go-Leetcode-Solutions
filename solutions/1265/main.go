package main

// Below is the interface for ImmutableListNode, which is already defined for you.
type ImmutableListNode struct {
}

func (this *ImmutableListNode) getNext() ImmutableListNode {
	// return the next node.
	return ImmutableListNode{}
}

func (this *ImmutableListNode) printValue() {
	// print the value of this node.
}

func printLinkedListInReverse(head ImmutableListNode) {
	// base case
	if head == nil {
		return
	}
	// dfs to the bottom first
	printLinkedListInReverse(head.getNext())
	// now print
	head.printValue()
}
