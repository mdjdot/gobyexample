package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./dat", d1, 0644)
	check(err)

	f, err := os.Create("./dat1")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Println(n2)
	fmt.Println(string(d2))

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Println(n3)

	f.Sync() // 调用 Sync 来将缓冲区的信息写入磁盘。

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n") // 9
	fmt.Println(n4)
	w.Flush() // 使用 Flush 来确保所有缓存的操作已写入底层写入器。
}
