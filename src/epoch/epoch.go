package main

import "fmt"
import "time"

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
$ go run src/epoch/epoch.go
2017-12-31 19:46:55.626353 +0800 CST m=+0.000370992
1514720815
1514720815626
1514720815626353000
2017-12-31 19:46:55 +0800 CST
2017-12-31 19:46:55.626353 +0800 CST
*/