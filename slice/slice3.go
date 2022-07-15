package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	fmt.Printf("len:%d, cap=%d, slice=%d\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素1, len+1, cap不变
	numbers = append(numbers, 1)
	fmt.Printf("len:%d, cap=%d, slice=%d\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len:%d, cap=%d, slice=%d\n", len(numbers), cap(numbers), numbers)

	// 如果容量cap已经满了, cap会扩容，容量为原来cap的一倍
	numbers = append(numbers, 1)
	fmt.Printf("len:%d, cap=%d, slice=%d\n", len(numbers), cap(numbers), numbers)

}
