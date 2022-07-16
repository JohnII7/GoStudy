package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine : i %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个go程 去执行newTask
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine : i %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
