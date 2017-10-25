package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 开启一个线程, 接收jobs发送的消息
	go func() {
		for {
			j, more := <-jobs // 监听jobs发送的消息，返回结果 j 代表消息，more代表是否接受了bool
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // 使用另一个channel 发送完毕true
				return
			}
		}
	}()

	// jobs channel 循环发送三次: 1, 2, 3
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	fmt.Println("sent all jobs")
	close(jobs) // 关闭jobs channel
	<-done // 在外检测另一个信道的发送内容
}

/*

// 由于存在线程问题, 输出顺序不是很确定
$ go run basic/closing-channels.go
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs

*/