package main

type MinStack struct {
	stack [][]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	// initialize our min stack with a dummy node, with current min to max int
	maxInt := int(^uint(0) >> 1)
	return MinStack{[][]int{{0, maxInt}}}
}

func (this *MinStack) Push(val int) {
	// push this value, and the min of this val, and current min
	newMin := min(val, this.GetMin())
	this.stack = append(this.stack, []int{val, newMin})
}

func (this *MinStack) Pop() {
	// pop right
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	// just last val in stack
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	// just last min in stack
	return this.stack[len(this.stack)-1][1]
}

// helper function for min
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
