package main

import "fmt"

func ping(pings chan <- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan <- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message") // 通过pings这个channel发送信息
	pong(pings, pongs) // 通过pongs发送(接收到的pings的消息)
	fmt.Println(<-pongs) // 接收pongs的消息
}

/*

$ go run Basic/channel-directions.go
passed message

*/
