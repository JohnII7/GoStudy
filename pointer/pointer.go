package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Println("after swap")
}

func main() {
	a := 10
	b := 20

	fmt.Println("a: ", a, ",b: ", b)
	// swap
	swap(&a, &b)
	fmt.Println("a: ", a, ",b: ", b)

	p := &a
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println("================")
	pp := &p
	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println(*pp)
}
