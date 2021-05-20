package main

type MovingAverage struct {
	queue []int
	sum   int
	size  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{[]int{}, 0, size}
}

func (this *MovingAverage) Next(val int) float64 {
	// add val to queue, and sum
	this.queue = append(this.queue, val)
	this.sum += val

	// check if queue exceeds size
	if len(this.queue) > this.size {
		// popleft, and take sum out
		first := this.queue[0]
		this.queue = this.queue[1:]
		this.sum -= first
	}

	// return sum / size of queue
	return float64(this.sum) / float64(len(this.queue))
}
