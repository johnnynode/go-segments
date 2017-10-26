package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile(p string) *os.File{
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// 写入文件
func writeFile(f * os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// 关闭文件
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}

/*

$ go run basic/defer.go
creating
writing
closing

*/

