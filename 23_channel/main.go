package main

import "fmt"

import "time"

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}
