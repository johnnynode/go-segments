package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one" // 发送第1条
	queue <- "two" // 发送第2条
	close(queue) // 此处close必不可少，关闭后for range迭代完成结束

	// 使用for range 迭代channel发送的信息
	for elem := range queue {
		fmt.Println(elem) // one ↵ two
	}
}
