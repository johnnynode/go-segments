package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	fmt.Println("one------>")

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	fmt.Println("two------>")

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	fmt.Println("thr------>")
}

/*
$ go run src/if-else/if-else.go
7 is odd
one------>
8 is divisible by 4
two------>
9 has 1 digit
thr------>
*/