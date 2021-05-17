package main

type CombinationIterator struct {
	combinations      []string
	characters        string
	combinationLength int
	index             int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	// get new instance, generate combos, return
	combinationIterator := CombinationIterator{[]string{}, characters, combinationLength, 0}
	combinationIterator.generate(&[]byte{}, 0)
	return combinationIterator
}

// helper function to recursively generate combinations
func (this *CombinationIterator) generate(temp *[]byte, start int) {
	// check if we have valid combo
	if len(*temp) == this.combinationLength {
		this.combinations = append(this.combinations, string(*temp))
	}

	// go through possible characters
	for i := start; i < len(this.characters); i++ {
		// take character, backtrack, remove
		*temp = append(*temp, this.characters[i])
		this.generate(temp, i+1)
		*temp = (*temp)[:len(*temp)-1]
	}
}

func (this *CombinationIterator) Next() string {
	// get current
	this.index++
	return this.combinations[this.index-1]
}

func (this *CombinationIterator) HasNext() bool {
	// check if index is less than combos
	return this.index < len(this.combinations)
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
