package main

import (
	"fmt"
	"time"
)

func conv() {

	var ch1 = make(chan int) // Канал для оригинальных чисел
	var ch2 = make(chan int) // Канал для их квадратов

	// Горутина для внесения оригинальных числе в первый канал
	go func() {
		for i := 0; ; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	// Горутина для чтения чисел из первого канала и возведения их в квадрат
	go func() {
		for {
			n := <-ch1
			n *= n
			ch2 <- n
			time.Sleep(1 * time.Second)
		}
	}()

	// Горутина для вывода в консоль
	go func() {
		for {
			fmt.Println(<-ch2)
			time.Sleep(1 * time.Second)
		}
	}()

	select {} // Ожидание окончания горутин

}
