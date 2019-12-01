package main

import "os"

import "fmt"

func main() {
	// panic("a problem")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("catch panic:", r)
		}
	}()
	_, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
	}
}
