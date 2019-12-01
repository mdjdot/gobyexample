package main

import "fmt"
import s "strings"

func main() {
	var p = fmt.Println
	p("contains:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("hasprefix:", s.HasPrefix("test", "te"))
	p("hassuffix:", s.HasSuffix("test", "st"))
	p("index:", s.Index("test", "kes")) // -1
	p("join:", s.Join([]string{"this", "is", "a", "string", "slice"}, "-"))
	p("repeat:", s.Repeat("apple", 3))                                           // count 0 返回空字符串
	p("replace:", s.Replace("an apple pie, i like apple", "apple", "banana", 1)) // replace: an banana pie, i like apple
	p("split:", s.Split("this is an apple", " "))
	p("tolower:", s.ToLower("UPPER STRING"))
	p("toupper:", s.ToUpper("lower string"))
	p()
}
