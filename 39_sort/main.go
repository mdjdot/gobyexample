package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	fmt.Println(strs)

	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{7, 2, 5}
	fmt.Println(ints)

	sort.Ints(ints)
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) // 倒序
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)
}
