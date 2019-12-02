package main

import "encoding/base64"

import "fmt"

func main() {
	data := "abc23!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // YWJjMjMhPyQqJigpJy09QH4=
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(sDec)         // [97 98 99 50 51 33 63 36 42 38 40 41 39 45 61 64 126]
	fmt.Println(string(sDec)) // abc23!?$*&()'-=@~

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc) // YWJjMjMhPyQqJigpJy09QH4=
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(uDec)         // [97 98 99 50 51 33 63 36 42 38 40 41 39 45 61 64 126]
	fmt.Println(string(uDec)) // abc23!?$*&()'-=@~
}

/*
标准base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），但是两者都可以正确解码为原始字符串。
*/
