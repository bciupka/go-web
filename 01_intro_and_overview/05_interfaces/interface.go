package main

import "fmt"

type IAnimal interface {
	Says() string
	CountLegs() int
}

type Dog struct {
	Name string
	Age  int
}

type Cat struct {
	Name           string
	WhishersLength int
}

func (*Dog) Says() string {
	return "Whoof"
}

func (*Cat) Says() string {
	return "Miau"
}

func (*Dog) CountLegs() int {
	return 4
}

func (*Cat) CountLegs() int {
	return 4
}

func AnimalInfo(a IAnimal) {
	fmt.Println("Animal says", a.Says(), "and has", a.CountLegs(), "legs.")
}
func main() {
	dog := Dog{
		"Torin",
		6,
	}

	cat := Cat{
		"Enzo",
		15,
	}

	AnimalInfo(&dog)
	AnimalInfo(&cat)
}
