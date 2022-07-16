package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		// 用go创建承载一个形参为空, 返回值为空的一个函数
		go func() {
			defer fmt.Println("A.defer")
			func() {
				defer fmt.Println("B.defer")

				// 立即终止当前goroutine(外层也会终止)
				runtime.Goexit()

				// 退出当前goroutine
				fmt.Println("B")
			}()
			fmt.Println("A")
		}() //尾部括号表示自执行函数，定义后即刻使用
	*/
	go func(a int, b int) bool {
		fmt.Println("a:", a, " b:", b)
		return true
	}(10, 20)
	for {
		time.Sleep(1 * time.Second)
	}
}
