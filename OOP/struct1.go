package main

import "fmt"

// 声明一种新的数据类型, 类型为int, 别名myint
type myint int

// Book 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
	book.auth = "666"
}
func changeBook2(book *Book) {
	book.auth = "666"
}
func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "John"

	fmt.Println(book1)
	changeBook1(book1)
	fmt.Println(book1)
	changeBook2(&book1)
	fmt.Println(book1)

}
