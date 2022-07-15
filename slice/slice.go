package main

import "fmt"

// 切片是引用传递
func printArraySlice(myArray []int) {
	// _下划线表示匿名变量
	for _, value := range myArray {
		fmt.Println("value:", value)
	}
}

func main() {
	// 动态数组,即slice切片
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray type:%T\n", myArray)

	printArraySlice(myArray)

}
