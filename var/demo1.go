package main

import "fmt"

/*四种声明方式*/

func main() {
	var a int
	fmt.Println("a: ", a)

	var b = 10
	fmt.Println("b: ", b)
	fmt.Printf("typeof b: %T\n", b)

	var c = 100
	fmt.Println("c: ", c)

	// 此方法不支持全局变量
	d := 1000
	fmt.Println("d: ", d)

}
