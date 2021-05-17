package main

type RecentCounter struct {
	stream []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	// push this to stream
	this.stream = append(this.stream, t)

	// keep removing from stream, as so long as its older than 300
	for this.stream[0] < t-3000 {
		this.stream = this.stream[1:]
	}

	return len(this.stream)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
