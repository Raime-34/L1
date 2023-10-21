package main

import "fmt"

func makeSet() {

	stringsSlice := []string{"apple", "banana", "cherry", "date", "banana", "fig", "grape", "cherry"}
	set := make(map[string]struct{}) // Map с пустыми структурами, ключ сохраняет свойство уникальности, так что можно его заобузить

	for _, string := range stringsSlice {
		set[string] = struct{}{}
	}

	for string, _ := range set {
		fmt.Print(string, " ")
	}

}
