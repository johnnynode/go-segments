package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1) // 睡眠1s后
		c1 <- "one" // c1 发送消息
	}()

	go func() {
		time.Sleep(time.Second * 2) // 睡眠2s后
		c2 <- "two" // c2 发送消息
	}()

	// 循环2次select 在select 中接收消息
	for i := 0; i < 2; i ++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/*

$ go run Basic/select.go
received one
received two

*/