package main

import "fmt"

import "time"

func main() {
	i := 2
	fmt.Print("write ", i, " as\n")
	fmt.Println("write", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	n := 5
	switch {
	case n > 0:
		fmt.Println("n is grater than 0")
		fallthrough
	case n > 5:
		fmt.Println("n is grater than 5")
		fallthrough
	case n > 10:
		fmt.Println("n is grater than 10")
		fallthrough
	default:
		fmt.Println("check value done")
	}
}

/*
fallthrough 从验证通过的分支开始执行之后的分支，不管之后是否通过
*/
