package main

import "fmt"

// interface{}是万能数据类型, 与interface不同, 多了一个花括号
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	// 给interface{}提供断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value:", value)
		fmt.Printf("value type:%T", value)
	}
}

type Book1 struct {
	auth string
}

func main() {
	book := Book1{"Golang"}
	myFunc(book)
	myFunc("Hello")
}
