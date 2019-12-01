package main

import "regexp"

import "fmt"

import "bytes"

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))                              // true
	fmt.Println(r.FindString("cde peach punch abcd"))                // peach
	fmt.Println(r.FindStringIndex("cde peach punch abcd"))           // [4 9]
	fmt.Println(r.FindStringSubmatch("cde peach punch abcd"))        // [peach ea]
	fmt.Println(r.FindStringSubmatchIndex("cde peach punch abcd"))   // [4 9 5 7]
	fmt.Println(r.FindAllString("cde peach punch abcd", -1))         // [peach punch]
	fmt.Println(r.FindAllStringSubmatch("cde peach punch abcd", -1)) // [[peach ea] [punch un]]
	fmt.Println(r.Match([]byte("cde peach punch abcd")))             // true

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // p([a-z]+)ch

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}
