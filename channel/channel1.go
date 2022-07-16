package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine正在运行...")

		// 将666发送给c
		c <- 666
	}()
	// 从channel中取出666
	num := <-c
	fmt.Println("num:", num)
	fmt.Println("main goroutine结束")
}
