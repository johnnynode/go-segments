package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping" // 发送
	}()

	msg := <-messages // 接收
	fmt.Printf(msg)
}
