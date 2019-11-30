package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Printf("7.0/3.0=%.2f\n", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
