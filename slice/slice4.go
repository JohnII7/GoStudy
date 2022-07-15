package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// [0,2)
	s1 := s[0:2]
	fmt.Println("s1:", s1)

	// s,s1底层数组是一样的，改变其中一个，其他的也会改变
	s1[0] = 100
	fmt.Println("s:", s, " s1:", s1)

	// s2=[0,0,0]
	// copy则不会发生上述情况
	s2 := make([]int, 3)
	copy(s2, s)
	s2[0] = 2222
	fmt.Println("s2:", s2, " s", s)
}
