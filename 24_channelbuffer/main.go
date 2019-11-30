package main

import "fmt"

func main() {
	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"
	// message <- "anotherthing" // 接受满了，再接受造成deadlock

	fmt.Println(<-message) // buffered 先进先出
	fmt.Println(<-message)
}
