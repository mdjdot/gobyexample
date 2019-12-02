package main

import "fmt"

import "os"

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

/*
os.Exit将不会执行defer
*/
