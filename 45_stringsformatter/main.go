package main

import "fmt"

import "os"

type point struct {
	x, y int
}

func main() {
	p := point{x: 1, y: 2}
	fmt.Printf("%v\n", p)                            // {1 2}
	fmt.Printf("%+v\n", p)                           // {x:1 y:2}
	fmt.Printf("%#v\n", p)                           // main.point{x:1, y:2}
	fmt.Printf("%T\n", p)                            // main.point
	fmt.Printf("%t\n", true)                         // true
	fmt.Printf("%d\n", 12)                           // 12
	fmt.Printf("%b\n", 12)                           // 1100
	fmt.Printf("%c\n", 33)                           // !
	fmt.Printf("%x\n", 33)                           // 21
	fmt.Printf("%f\n", 33.3)                         // 33.300000
	fmt.Printf("%e\n", 33000.0)                      // 3.300000e+04
	fmt.Printf("%E\n", 33000.0)                      // 3.300000E+04
	fmt.Printf("%s\n", "\"string\"")                 // "string"
	fmt.Printf("%q\n", "\"string\"")                 // "\"string\""
	fmt.Printf("%x\n", "hex this")                   // 6865782074686973
	fmt.Printf("%p\n", &p)                           // 0xc0000120c0
	fmt.Printf("%6d%6d\n", 12, 345)                  //     12   345
	fmt.Printf("|%6.2f|%6.2f|\n", 12.123, 345.456)   // | 12.12|345.46|
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12.123, 345.456) // |12.12 |345.46|
	fmt.Printf("|%6s|%6s|\n", "abc", "de")           // |   abc|    de|
	fmt.Printf("|%-6s|%-6s|\n", "abc", "de")         // |abc   |de    |
	s := fmt.Sprintf("a %s", "string")               // a string
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error") // a string
}
