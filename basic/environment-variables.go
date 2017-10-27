package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	os.Setenv("FOO", "1") // 设置环境变量
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	os.Unsetenv("FOO") // 移除环境变量

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}