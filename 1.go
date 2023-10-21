package main

import "fmt"

type Human struct {
	fName string
	lName string
}

type Action struct {
	Human // "Наследование" свойств от Human
	aName string
}

// Реализация receiver function для структуры Human
func (h *Human) getName() string {
	return fmt.Sprintf("%v %v", h.fName, h.lName)
}
