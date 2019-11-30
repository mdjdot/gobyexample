package main

import "fmt"

// Person .
type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Bob", age: 20})
	fmt.Println(Person{
		name: "Bob",
		age:  20,
	})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40}) // &{Ann 40}

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
