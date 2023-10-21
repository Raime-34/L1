package main

import (
	"fmt"
	"sync"
	"time"
)

func fewfn() {

	var wg sync.WaitGroup
	// wg.Add(5)
	defer wg.Wait()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			isPastYourBedTime(i*2, i+1)
			wg.Done()
		}(i)
	}
}

// Сама функция для имитации работы sleep
func isPastYourBedTime(n int, i int) {

	var t = time.After(time.Duration(n) * time.Second) // Инициализация канала, в который через n секунд будет проброщено текущее время
	fmt.Printf("Горутина №%v спит\n", i)               //
	<-t                                                // Остановка текущей горутины на n секунд
	fmt.Printf("Горутина №%v проснулась\n", i)         //

}
