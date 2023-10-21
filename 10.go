package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func tempSeq() {
	list := make([]float64, 0)             // Массив, в который я закидываю сгенеренные температуры
	var m = make(map[float64][]float64, 0) // Map для хранения значений из массив сверху по ключю из ближайших чисел кратных 10, например m[-20] = { -23.3, -24.5, ... }
	tempGenerator(&list)                   // Генерация случайных значений температур

	// Группировки температур в Map
	for _, v := range list {
		closestTemp := roundToMultipleOf10(v)                          // Получение ближайщего наименьего кратного 10 числа
		m[closestTemp] = append(m[closestTemp], float64(int(v*10))/10) // Добавление в соответствующий слайс температуры, + я оставляю лишь один знак после запятой, чтобы это не выглядело грамостко в выводе
	}

	keys := make([]float64, 0) // Слайс для ключей, понадобиться чтобы в порядке возрастания температуры выводить значения
	for temp, _ := range m {
		keys = append(keys, temp)
	}
	sort.Float64s(keys) // Сама сориторвка слайса из ключей

	for _, key := range keys {
		fmt.Println(key, m[key]) // Вывод итогового результата
	}
}

// Просто генератор чисел
func tempGenerator(list *[]float64) {

	min := -64.9 // Верхняя и нижняя
	max := 102.0 // граница для генератора

	// Генерация с шагом от 0.0 до 10.0, правая граница не влючается
	for ; min < max; min += rand.Float64() * 10 {
		*list = append(*list, min)
	}

}

// Функция округления к ближайщему наименьшему кратному 10 числу
func roundToMultipleOf10(x float64) float64 {
	result := math.Floor(math.Abs(x/10.0)) * 10.0 // Беру модуль числа, делю на 10, округляю в меньшую сторону, обратно умножаю на 10, получаю ближайщее кратное 10 число
	if x < 0 {
		result *= -1
	}
	return result
}
