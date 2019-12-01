package main

import "time"

import "fmt"

func main() {
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(15 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

/*
定时器 是当你想要在未来某一刻执行一次时使用的。
打点器 则是当你想要在固定的时间间隔重复执行准备的。
*/
