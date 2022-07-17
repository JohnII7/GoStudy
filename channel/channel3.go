package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 关闭一个channel
		close(c)
	}()
	/*
		for {
			// 如果ok为true表示channel没有关闭，false表示channel已经关闭
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/
	// 也可以使用range来迭代操作channel
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("main finished")
}
