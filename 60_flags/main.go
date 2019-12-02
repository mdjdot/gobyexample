package main

import "flag"

import "fmt"

func main() {
	p := fmt.Println
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("num", 43, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()
	p("word", *wordPtr)
	p("num", *numPtr)
	p("fork", *boolPtr)
	p("svar", svar)
	p("tail", flag.Args()) // Args returns the non-flag command-line arguments.
}

/*
go run .\main.go -word="abc d" -num=23 -fork=true a1 b2 c3
word abc d
num 23
fork true
svar bar
tail [a1 b2 c3]
*/
