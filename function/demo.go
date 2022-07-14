package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 777
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 1000
	r2 = 2000
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 1000
	r2 = 2000
	return
}
func main() {
	c := foo1("abc", 12)
	fmt.Println("c = ", c)

	res1, res2 := foo2("asd", 12)
	fmt.Println("res1 = ", res1)
	fmt.Println("res2 = ", res2)

	res1, res2 = foo3("foo3", 333)
	fmt.Println("res1 = ", res1)
	fmt.Println("res2 = ", res2)

	res1, res2 = foo3("foo4", 333)
	fmt.Println("res1 = ", res1)
	fmt.Println("res2 = ", res2)

}
