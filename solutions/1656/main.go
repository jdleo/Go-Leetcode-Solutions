package main

type OrderedStream struct {
	stream []string // stream to hold strings
	ptr    int      // pointer for current string
}

func Constructor(n int) OrderedStream {
	// make stream array of size n+1
	return OrderedStream{make([]string, n+1), 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	res := []string{}
	// insert string into stream
	this.stream[idKey-1] = value

	// keep checking as long as we can pull strings at pointer
	for this.stream[this.ptr-1] != "" {
		// put this string into res, and increment ptr
		res = append(res, this.stream[this.ptr-1])
		this.ptr++
	}

	return res
}
