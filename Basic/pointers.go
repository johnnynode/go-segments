package main

import "fmt"

// 不会改变原值
func zeroval (ival int) {
	ival = 0
}

// 改变原值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i) // zeroval: 1

	zeroptr(&i)
	fmt.Println("zeroptr:", i) // zeroptr: 0

	fmt.Println("pointer:", &i) // pointer: 0xc0420461d0
}

/*

$ go run Basic/pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc0420461d0

*/
