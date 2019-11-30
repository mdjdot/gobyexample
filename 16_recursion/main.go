package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}

/*
f0 = 1
f1 = 1*f0
fn = n*f(n-1)
*/
