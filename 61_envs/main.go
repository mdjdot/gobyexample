package main

import "os"

import "fmt"

import "strings"

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO")) // FOO: 1
	fmt.Println("BAR:", os.Getenv("BAR")) // BAR:

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
