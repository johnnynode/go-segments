package main

import "time"
import "fmt"

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
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*
$ go run src/rate-limiting/rate-limiting.go
request 1 2017-12-31 10:19:06.169104 +0800 CST m=+0.201675582
request 2 2017-12-31 10:19:06.368149 +0800 CST m=+0.400714860
request 3 2017-12-31 10:19:06.571239 +0800 CST m=+0.603798064
request 4 2017-12-31 10:19:06.773037 +0800 CST m=+0.805590281
request 5 2017-12-31 10:19:06.972522 +0800 CST m=+1.005068615
request 1 2017-12-31 10:19:06.972651 +0800 CST m=+1.005198305
request 2 2017-12-31 10:19:06.972672 +0800 CST m=+1.005218882
request 3 2017-12-31 10:19:06.972684 +0800 CST m=+1.005231232
request 4 2017-12-31 10:19:07.174007 +0800 CST m=+1.206548312
request 5 2017-12-31 10:19:07.373543 +0800 CST m=+1.406078520
*/