package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<- done
	fmt.Println("exiting")
}

/*

$ go run signals.go
awaiting signal
^C
interrupt // 此处并不是输出内容，只是演示提示
exiting // 退出

*/