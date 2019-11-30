package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func plus(a int, b int) int {
	return a + b
}

func subs(a, b int) (result int) {
	result = a - b
	return result
}

func main() {
	sayHi()

	result := plus(1, 2)
	fmt.Println(result)

	result = subs(1, 3)
	fmt.Println(result)
}
