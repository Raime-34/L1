package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func insertIntoMap() {

	var m = make(map[int]interface{})
	var rwMutex sync.RWMutex       // Понадобится для предоставление доступ разных горутин к map
	var c = make(chan interface{}) // Понадобится для чтения данных, которые инсертяться в map
	// var signals = make(chan os.Signal, 1)
	// signal.Notify(signals, os.Interrupt)
	// var wg sync.WaitGroup
	// wg.Add(1)

	fmt.Println("Элементы map:")

	// Горутина для вывода новых элементов map
	go func() {
		for {
			val, ok := <-c
			if ok {
				fmt.Printf("[%v] ", val)
			}
		}
	}()

	// Горутина для инсерта в map текущего времени
	go func() {
		for {
			rwMutex.Lock()                  // Если мютекс не заблокирован,
			newElement := time.Now().Unix() // то горутина выполняет следующие три команды,
			c <- newElement                 // в ином случае ожидает пока мютекс будт раблокирован
			m[len(m)-1] = newElement        //
			rwMutex.Unlock()                // Разблокировка мютекса
			time.Sleep(1 * time.Second)
		}
	}()

	// Горутина для инсерта случайной строки
	go func() {
		for {
			rwMutex.Lock()
			newElement := randSeq(7)
			c <- newElement
			m[len(m)-1] = newElement
			rwMutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	// go func() {
	// 	for {
	// 		<-signals
	// 		wg.Done()
	// 		return
	// 	}
	// }()

	select {} // Позволяет приостановить текущий поток, пока выполняются горутины
	// wg.Wait()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
