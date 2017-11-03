package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // 接收的channel有两个返回值,一个是信息,一个是是否含有更多
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // 发送true
				return
			}
		}
	}()

	// 循环3次发送jobs的channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // 关闭jobs这个channel
	fmt.Println("sent all jobs")
}

/*

// 可能的输出结果
$ go run Basic/closing-channels.go
sent job 1
sent job 2
sent job 3
sent all jobs

// 可能的输出结果
$ go run Basic/closing-channels.go
sent job 1
sent job 2
received job 1
received job 2
received job 3
sent job 3
sent all jobs

*/