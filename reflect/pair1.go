package main

import "fmt"

func main() {
	var a string
	a = "John"

	var allType interface{}
	allType = a
	fmt.Println(allType)

	str, _ := allType.(string)
	fmt.Println(str)
}
