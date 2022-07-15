package main

import "fmt"

// 数组是值传递
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index:", index, " value:", value)
	}
}

func main() {
	// 固定长度的数组
	var myArray1 [2]int

	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}
	for index, value := range myArray2 {
		fmt.Println("index:", index, " value:", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types =%T\n", myArray1)
	fmt.Printf("myArray2 types =%T\n", myArray2)

	// 数组传参
	printArray(myArray3)
}
