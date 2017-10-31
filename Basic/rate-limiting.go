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

	for i := 0; i < 3; i ++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*

$ go run Basic/rate-limiting.go
request 1 2017-10-31 14:03:48.7719383 +0800 CST
request 2 2017-10-31 14:03:48.9725747 +0800 CST
request 3 2017-10-31 14:03:49.1721922 +0800 CST
request 4 2017-10-31 14:03:49.3718522 +0800 CST
request 5 2017-10-31 14:03:49.5726871 +0800 CST
request 1 2017-10-31 14:03:49.5726871 +0800 CST
request 2 2017-10-31 14:03:49.5726871 +0800 CST
request 3 2017-10-31 14:03:49.5726871 +0800 CST
request 4 2017-10-31 14:03:49.7733573 +0800 CST
request 5 2017-10-31 14:03:49.9731115 +0800 CST

*/
