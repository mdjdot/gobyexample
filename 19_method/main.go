package main

import "fmt"

// Rect .
type Rect struct {
	width, height int
}

func (r *Rect) area() int {
	return r.width * r.height
}

func (r *Rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := Rect{width: 10, height: 5}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
