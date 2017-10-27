package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a) // 3
	fmt.Println(b) // 7

	_, c := vals()
	fmt.Println(c) // 7
}
