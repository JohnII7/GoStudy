package main

import "fmt"

func main() {
	// ===>第一种声明方式
	// map类型, key:string, value:string
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap是一个空map")
	}
	// 在使用map前, 需要先用make给map分配数据空间
	myMap = make(map[string]string, 10)

	myMap["one"] = "java"
	myMap["two"] = "go"
	myMap["three"] = "python"
	fmt.Println("myMap1:", myMap)

	// ===>第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "rust"
	fmt.Println("myMap2:", myMap2)

	// ===>第三种声明方式
	myMap3 := map[string]string{
		"1": "John",
		"2": "Riki",
		"3": "ChungYui",
	}
	fmt.Println("myMap3:", myMap3)
}
