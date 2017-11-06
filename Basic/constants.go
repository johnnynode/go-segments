package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000
	const m = 5e8
	const q = 5E8

	fmt.Println(n == m)
	fmt.Println(m == q)

	const d = 3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

/*
$ go run Basic/constants.go
constant
true
true
6e+11
600000000000
-0.28470407323754404

*/
