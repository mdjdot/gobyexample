package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	println(sum)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
