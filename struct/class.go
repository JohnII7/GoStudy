package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) Show() {
	//fmt.Println("hero:", this)
	fmt.Println("name:", this.Name)
	fmt.Println("ad:", this.Ad)
	fmt.Println("level:", this.Level)
}
func (this *Hero) GetName() string {
	fmt.Println("Name:", this.Name)
	return this.Name
}
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "John", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("RIKI")
	hero.Show()
}
