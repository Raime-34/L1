package main

import "fmt"

// Удаление без сохранения оригинальной последовательности
func removeAt1() {

	slice := []int{1, 2, 3, 4, 5}
	index := 1 // Индекс удаляемого элемнта

	if index < len(slice) {
		slice[index] = slice[len(slice)-1]   // Перемещаем последний элемент на место того, которого хотим удалить
		slice = append(slice[:len(slice)-1]) // Удаляем последний элемент так как он по сути теперь дублирует slice[i]
		fmt.Println(slice)
	}

}

// Удаление с сохранением порядка
func removeAt2() {

	slice := []int{1, 2, 3, 4, 5}
	index := 1 // Индекс удаляемого элемнта

	if index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
		fmt.Println(slice)
	}

}
