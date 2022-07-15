package main

import "fmt"

// defer和return的顺序
// return先执行, defer后执行
func deferFunc() int {
	fmt.Println("defer called")
	return 0
}

func returnFunc() int {
	fmt.Println("return called")
	return 0

}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
