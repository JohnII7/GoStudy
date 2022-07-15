package lib1

import "fmt"

// 当前lib包提供的api
func Lib1Test() {
	fmt.Println("Lib1Test()...")
}

func init() {
	fmt.Println("lib1, init()...")
}
