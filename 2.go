package main

import (
	"fmt"
	"sync"
)

// Варинат с каналом
func makeItSquared1(num []int) (result []int) {

	result = make([]int, len(num))
	var ch = make(chan bool)
	defer close(ch)

	// Запуск для кажого элемента слайса горутины
	for index, value := range num {
		go func(index int, value int) { // Передача копии итератора, для корректной работы анонимной функции
			result[index] = value * value
			ch <- true // Подтверждение завершения работы горутины
		}(index, value)
	}

	// Ожидание получения len(num) значений из канала, для того чтобы гарантировать завершение всех горутин
	for i := 0; i < len(num); i++ {
		<-ch
	}

	fmt.Println(result)
	return
}

// Вариант с WaitGroup, код почти идентичный
func makeItSquared2(num []int) (result []int) {

	result = make([]int, len(num))
	var wg sync.WaitGroup
	wg.Add(len(num))

	for index, value := range num {
		go func(index int, value int) {
			result[index] = value * value
			wg.Done()
		}(index, value)
	}

	wg.Wait()
	fmt.Println(result)
	return

}
