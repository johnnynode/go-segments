package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

/*

$ go run Basic/epoch.go
2017-11-05 05:39:09.154715 +0800 CST
1509831549
1509831549154
1509831549154715000
2017-11-05 05:39:09 +0800 CST
2017-11-05 05:39:09.154715 +0800 CST

*/
