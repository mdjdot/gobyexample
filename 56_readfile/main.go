package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./dat") // 调用os.Open和f.Close()
	check(err)
	fmt.Println(dat)         // [104 101 108 108 111 32 119 111 114 108 100 13 10]
	fmt.Println(string(dat)) // hello world

	f, err := os.Open("./dat")
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Println(n1)         // 5
	fmt.Println(string(b1)) // hello

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Println(n2)         // 2
	fmt.Println(o2)         // 6
	fmt.Println(string(b2)) // wo

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Println(n3)         // 2
	fmt.Println(o3)         // 6
	fmt.Println(string(b3)) // wo

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Println(b4)         // [114 108 100 13 10]
	fmt.Println(string(b4)) // rld

}
