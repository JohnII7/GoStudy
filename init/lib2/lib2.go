package lib2

import "fmt"

// Lib2Test 当前lib包提供的api
func Lib2Test() {
	fmt.Println("Lib2Test()...")
}

func init() {
	fmt.Println("lib2, init()...")
}
