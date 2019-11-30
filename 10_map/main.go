package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)
	fmt.Println(m["k2"])
	fmt.Println(m["kn"]) // 默认零值
	fmt.Println(len(m))  // 查询不存在的元素，长度不变

	v1 := m["k1"]
	fmt.Println(v1)

	delete(m, "k2")
	fmt.Println(m)
	fmt.Println(m["k2"])
	fmt.Println(len(m))

	n := map[string]int{"k1": 10, "k2": 21}
	fmt.Println(n)

	o := map[string]int{
		"k1": 21,
		"k2": 31,
	}
	fmt.Println(o)
}
