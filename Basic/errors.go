package main

import (
	"errors"
	"fmt"
)

// f1 函数如果参数为42 返回-1，errorMessage; 否则返回参数+3，nil
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42") // errors的api
	}
	return arg + 3, nil
}

// 创建一个argError 的struct
type argError struct {
	arg  int
	prob string
}

// 给argError 创建一个Error方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// f2函数如果参数为42,则返回-1,错误信息；否则返回参数+3，nil
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e) // f1 failed: can't work with 42
		} else {
			fmt.Println("f1 worked:", r) // f1 worked: 10
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e) // f2 failed: 42 - "can't work with it"}
		} else {
			fmt.Println("f2 worked:", r) // f2 worked: 10
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg) // 42
		fmt.Println(ae.prob) // can't work with it
	}
}

/*

$ go run Basic/errors.go
f1 worked: 10
f1 failed: can't work with 42
f2 worked: 10
f2 failed: 42 - can't work with it
42
can't work with it

*/