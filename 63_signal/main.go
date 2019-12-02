package main

import "os"

import "os/signal"

import "syscall"

import "fmt"

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/*
有时候，我们希望 Go 能智能的处理 Unix 信号。
例如，我们希望当服务器接收到一个 SIGTERM 信号时能够自动关机，
或者一个命令行工具在接收到一个 SIGINT 信号时停止处理输入信息。
这里讲的就就是在 Go 中如何通过通道来处理信号。
*/
