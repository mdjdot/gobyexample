package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Printf("%d: %d\n", i, num)
		}
	}

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range kvs {
		fmt.Printf("%s: %s", k, v)
	}

	for i, c := range "who am i" {
		// fmt.Println(i, c) // c rune
		fmt.Println(i, string(c))
	}
}
