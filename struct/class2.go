package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat")
}
func (this *Human) Walk() {
	fmt.Println("Human walk")
}

type SuperMan struct {
	Human // 表示SuperMan继承了Human
	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan eat")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan fly")
}

func main() {
	h := Human{"John", "male"}
	h.Eat()
	h.Walk()

	// 定义子类对象
	s := SuperMan{Human{"Wreck", "24"}, 100}
	s.Walk()
	s.Eat()
	s.Fly()
	fmt.Println(s.name)
}
