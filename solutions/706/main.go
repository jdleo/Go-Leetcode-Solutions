package main

type MyHashMap struct {
	array [1_000_001]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	array := [1_000_001]int{}
	for i := 0; i < 1_000_001; i++ {
		array[i] = -1
	}
	return MyHashMap{array}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.array[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.array[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.array[key] = -1
}
