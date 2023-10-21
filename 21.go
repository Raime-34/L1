package main

import "fmt"

// Интерфейс, который объявляет общую функцию для своих наследников
type Animal interface {
	makeNoises()
}

type Cat struct{}

// Фкнкция, которая вызывается при вызове makeNoises для catAdapter
func (it *Cat) meowing() {
	fmt.Println("Meow!")
}

// Сам адаптер
type CatAdapter struct {
	cat *Cat
}

// Определение функции из интерфейса Animal
func (adapter *CatAdapter) makeNoises() {
	adapter.cat.meowing()
}

type Dog struct{}

func (it *Dog) barking() {
	fmt.Println("Woof!")
}

type DogAdapter struct {
	dog *Dog
}

func (adapter *DogAdapter) makeNoises() {
	adapter.dog.barking()
}

type Human struct{}

func (it *Human) SaySomething() {
	fmt.Println("ACAB")
}

type HumanAdapter struct {
	human *Human
}

func (adapter *HumanAdapter) makeNoises() {
	adapter.human.SaySomething()
}

func main() {

	cat := CatAdapter{cat: new(Cat)}
	dog := DogAdapter{dog: new(Dog)}
	FloridaMan := HumanAdapter{human: new(Human)}

	cat.makeNoises()
	dog.makeNoises()
	FloridaMan.makeNoises()

}
