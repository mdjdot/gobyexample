package main

import "fmt"

func main() {
	s := make([]string, 3)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	s = append(s, []string{"g", "h", "i"}...)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	l := s[:5]
	fmt.Println(l)
	fmt.Println(len(l))
	fmt.Println(cap(l))

	l = s[2:]
	fmt.Println(l)
	fmt.Println(len(l))
	fmt.Println(cap(l))

	t := []string{"j", "k", "l"}
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1) //  多维slice内部slice长度可以不一样
		for j := 0; j < i+1; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
