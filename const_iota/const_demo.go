package main

import "fmt"

// 定义枚举
// 第一行的iota为0，iota会累加；若第一行为iota*10，则每次累加10
// iota只能配合const使用
const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("beijing: ", BEIJING)
	fmt.Println("shanghai: ", SHANGHAI)
	fmt.Println("shenzhen: ", SHENZHEN)
}
