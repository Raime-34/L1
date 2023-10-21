package main

import (
	"fmt"
)

func aaaaaaa() {
	a := []int{68, 37, 44, 15, 26, 52, 62, 86, 85, 70, 67, 35, 27, 38, 40, 46, 71, 24, 23, 49,
		12, 81, 41, 84, 33, 10, 28, 9, 69, 17, 79, 42, 64, 1, 31, 60, 29, 91, 3, 87, 94,
		65, 77, 5, 51, 32, 7, 80, 59, 74, 48, 88, 36, 13, 2, 19, 50, 34, 78, 92, 18, 73,
		21, 61, 63, 6, 14, 8, 58, 72, 66, 83, 4, 43, 53, 45, 82, 22, 54, 20, 30, 56, 90,
		25, 55, 76, 75, 95, 89, 16, 0, 57, 11, 47, 99, 98, 93, 96, 97}

	quicksort(a, 0, len(a)-1)

	fmt.Println(a)
}

// Решение основано на https://www.youtube.com/watch?v=Hoixgm4-P4M
func quicksort(slice []int, left int, right int) {

	orLeft := left // Заоминаю оригинальную левую граину, для дальнейшего рекурсиовного вызова

	// Проверка не передан ли был пустой кусок массива
	if left >= right {
		return
	}

	midle := (left + right + 1) / 2                         // Высчитывание индекса центрального элемента в текущей итерации рекурсии
	slice[midle], slice[right] = slice[right], slice[midle] // Меняем местами последний и центральныйэлемент

	// В следующем цикле left всегда указыват на элемент, который больше slice[right], когда находим i-ый элемент меньший или равный slice[right], меняемся с числом, большим slice[right]
	for i := left; i < right; i++ {
		if slice[i] <= slice[right] {
			slice[i], slice[left] = slice[left], slice[i]
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left] // Ставим ранее центральный элемент на его место

	quicksort(slice, orLeft, left-1)
	quicksort(slice, left+1, right)

}
