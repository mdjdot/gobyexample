package main

import "fmt"

func main() {
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		i++
		if i > 10 {
			break
		}
	}

	for k, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		fmt.Println(k, v)
	}
}
