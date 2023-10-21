package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Использование эксперементальной библиотеки slices, в данном случае тайм комплексити O(len(slice1) * len(slice2))
func intersection1() {

	slice1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	slice2 := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}

	fmt.Print("Пересечение: [ ")

	for _, v := range slice1 {
		if slices.Contains(slice2, v) {
			fmt.Print(v, " ")
		}
	}

	fmt.Print("]")
}
