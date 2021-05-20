package main

import (
	"strings"
)

type FileSystem struct {
	Value    int
	Children map[string]FileSystem
}

func Constructor() FileSystem {
	return FileSystem{-1, map[string]FileSystem{}}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	fullPath := strings.Split(path, "/")
	cur := *this

	for i := 1; i < len(fullPath)-1; i++ {
		// check if current level is in this children
		if v, ok := cur.Children[fullPath[i]]; ok {
			cur = v
		} else {
			// can't create
			return false
		}
	}

	// no issues, and create it, if it doesn't exist
	if _, ok := cur.Children[fullPath[len(fullPath)-1]]; ok {
		return false
	} else {
		cur.Children[fullPath[len(fullPath)-1]] = FileSystem{value, map[string]FileSystem{}}
		return true
	}
}

func (this *FileSystem) Get(path string) int {
	fullPath := strings.Split(path, "/")
	cur := *this

	// traverse through path
	for i := 1; i < len(fullPath); i++ {
		// check if current level is in this children
		if v, ok := cur.Children[fullPath[i]]; ok {
			cur = v
		} else {
			// not found
			return -1
		}
	}

	// return cur's value
	return cur.Value
}
