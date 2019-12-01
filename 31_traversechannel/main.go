package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // 通道不关就遍历会deadlock

	for element := range queue {
		fmt.Println(element)
	}

	// for {
	// 	if element, ok := <-queue; ok { // 不做ok判断，会一直输出false
	// 		fmt.Println(element, ok)
	// 	}
	// }
}
