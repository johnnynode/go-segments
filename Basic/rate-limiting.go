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
	close(burstyLimister)
	for req := range burstyRequests {
		<-burstyLimister
		fmt.Println("request", req, time.Now())
	}
}
