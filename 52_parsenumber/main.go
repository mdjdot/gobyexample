package main

import "strconv"

import "fmt"

func main() {
	f, _ := strconv.ParseFloat("1.234", 64) // ParseFloat converts the string s to a floating-point number with the precision specified by bitSize: 32 for float32, or 64 for float64.
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 2) // ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) // 456
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wait")
	if e != nil {
		fmt.Println(e)
	}

	s := strconv.Itoa(123)
	fmt.Println(s)
}
