package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1: ", v1) // v1: 7

	fmt.Println("len:", len(m)) // len: 2

	delete(m, "k2")
	fmt.Println("map:", m) // map: map[k1:7]

	_, prs := m["k2"]
	fmt.Println("prs:", prs) // prs: false

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n) // map: map[bar:2 foo:1]
}