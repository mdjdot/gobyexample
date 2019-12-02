package main

import "os"

import "fmt"

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)    // [xxx\main.exe 123 456 789]
	fmt.Println(argsWithoutProg) // [123 456 789]
	fmt.Println(arg)             // 789
}
