package main

import "fmt"

type rect struct {
	width, height int
}

// 求矩形面积函数
func (r *rect) area() int {
	return r.width * r.height
}

// 求矩形周长函数
func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{width:10, height:5} // 初始化一个矩形对象

	fmt.Println("area: ", r.area()) // area: 50
	fmt.Println("perim:", r.perim()) // perim: 30
	fmt.Println()

	rp := &r
	fmt.Println("area: ", rp.area()) // area: 50
	fmt.Println("preim:", rp.perim()) // preim: 30
	fmt.Println()

	rb := r
	fmt.Println("area: ", rb.area()) // area: 50
	fmt.Println("preim:", rb.perim()) // preim: 30
}

/*

$ go run Basic/methods.go
area:  50
perim: 30

area:  50
preim: 30

area:  50
preim: 30

*/

