package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	bsa := h.Sum([]byte("salt"))
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Printf("%x\n", bsa)

	s1 := "md5 this string"
	m := md5.New()
	m.Write([]byte(s1))
	bs1 := m.Sum(nil)
	bsa1 := m.Sum([]byte("salt"))
	fmt.Println(s1)
	fmt.Printf("%x\n", bs1)
	fmt.Printf("%x\n", bsa1)
}

/*
计算 MD5 散列，引入 crypto/md5 并使用 md5.New()方法。
sha1更安全，md5更快
*/
