package main

import "fmt"

import "time"

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(3 * time.Second)
	fmt.Println("done")

	done <- true // false/true都可以，只要有值
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
	fmt.Println("all done")
}

/*
通过接收到channel中的值，确认channel获得值之前的工作已完成
*/
