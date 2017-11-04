package main

import (
	"time"
	"fmt"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimister := make(chan time.Time, 3)

	for i := 0; i < 3; i ++ {
		burstyLimister <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimister <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimister
		fmt.Println("request", req, time.Now())
	}
}

/*

$ go run Basic/rate-limiting.go
request 1 2017-11-04 17:22:51.3472072 +0800 CST
request 2 2017-11-04 17:22:51.5488542 +0800 CST
request 3 2017-11-04 17:22:51.7475981 +0800 CST
request 4 2017-11-04 17:22:51.9473478 +0800 CST
request 5 2017-11-04 17:22:52.147337 +0800 CST
request 1 2017-11-04 17:22:52.147337 +0800 CST
request 2 2017-11-04 17:22:52.147337 +0800 CST
request 3 2017-11-04 17:22:52.147337 +0800 CST
request 4 2017-11-04 17:22:52.3480068 +0800 CST
request 5 2017-11-04 17:22:52.5481853 +0800 CST

*/
