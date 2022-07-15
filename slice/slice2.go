package main

import "fmt"

func main() {
	// 1.声明并且初始化
	//slice1 := []int{1, 2, 3}

	// 2.声明但未初始化
	//var slice1 []int
	// 开辟3个空间，初始化值为0
	//slice1 = make([]int, 3)

	// 3.声明，同时3个分配空间，初始化值为0
	slice1 := make([]int, 3)
	fmt.Printf("len:%d, slice:%d", len(slice1), slice1)

	// 判断一个slice是否为0
	if slice1 == nil {
		fmt.Println("slice1是一个空切片")
	} else {
		fmt.Println("slice 有空间")
	}
}
