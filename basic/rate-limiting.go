package main

import (
	"time"
	"fmt"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i;
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*

// 运行结果示例：
$ go run basic/rate-limiting.go
request 1 2017-10-25 17:19:52.1734175 +0800 CST
request 2 2017-10-25 17:19:52.3736813 +0800 CST
request 3 2017-10-25 17:19:52.5740786 +0800 CST
request 4 2017-10-25 17:19:52.7741063 +0800 CST
request 5 2017-10-25 17:19:52.9735633 +0800 CST
request 1 2017-10-25 17:19:52.9735633 +0800 CST
request 2 2017-10-25 17:19:52.9735633 +0800 CST
request 3 2017-10-25 17:19:52.9735633 +0800 CST
request 4 2017-10-25 17:19:53.1750039 +0800 CST
request 5 2017-10-25 17:19:53.3742565 +0800 CST

*/