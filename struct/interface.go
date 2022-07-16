package main

import "fmt"

// AnimalI 本质是一个指针
type AnimalI interface {
	Sleep()
	GetColor() string
	GetType() string
}

// Cat 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}
func showAnimal(animal AnimalI) {
	animal.Sleep()
	fmt.Println("color:", animal.GetColor())
	fmt.Println("kind", animal.GetType())
}
func main() {
	/*
		var animal AnimalI
		animal = &Cat{"Green"}
		animal.Sleep()

		animal = &Dog{"Red"}
		animal.Sleep()
	*/
	cat := Cat{"Green"}
	dog := Dog{"Red"}
	showAnimal(&cat)
	showAnimal(&dog)
}
