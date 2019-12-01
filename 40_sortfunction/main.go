package main

import (
	"fmt"
	"sort"
)

// ByLength .
type ByLength []string

// Len .
func (s ByLength) Len() int {
	return len(s)
}

// Swap .
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits)) // 实现sort.Interface，才能被sort.Sort使用
	fmt.Println(fruits)
}

/*
sort.Interface
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
