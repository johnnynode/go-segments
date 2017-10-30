package main

import "fmt"

func main() {
	messages := make(chan string) // 使用make new 出一个名为messages的chanel

	go func() {
		messages <- "ping" // 通过channel <-的语法发送一个值：ping
	}()

	msg := <-messages // 通过 <-channel的语法接收之前发送的值
	fmt.Println(msg) // ping
}

/*

$ go run Basic/channels.go
ping

*/