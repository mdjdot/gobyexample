package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if _, err := os.Stat("./dat"); err != nil {
		ioutil.WriteFile("./dat", []byte("hello\nfilter\n"), 777)
	}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}

/*
cat .\58_filter\dat | go run .\58_filter\main.go
HELLO
FILTER
*/
