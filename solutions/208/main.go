package main

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	// return new trie instance
	return Trie{root: &TrieNode{map[rune]*TrieNode{}, false}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root

	// go through word
	for _, c := range word {
		// check if not already in children
		if _, ok := cur.children[c]; !ok {
			// create
			cur.children[c] = &TrieNode{map[rune]*TrieNode{}, false}
		}
		// set next
		cur = cur.children[c]
	}

	// mark as end of word
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root

	// go through word
	for _, c := range word {
		// check if in children
		if child, ok := cur.children[c]; ok {
			// set cur
			cur = child
		} else {
			return false
		}
	}

	// check if end of word
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root

	// go through prefix
	for _, c := range prefix {
		// check if in children
		if child, ok := cur.children[c]; ok {
			// set cur
			cur = child
		} else {
			return false
		}
	}

	// no probs
	return true
}
