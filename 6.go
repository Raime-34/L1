package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Завершение с помощью закрытия канала
func withChanel1() {
	var c = make(chan int)
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)

	go func() {
		for {
			select {
			case <-c:
				fmt.Println("\tЗавершение с помощью закрытия канала")
				wg.Done()
				return
			default:
				fmt.Println("Heartbeat...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(c)
}

// Тотже код только кэйс <-c тригерит не закрытие канала, а проброска значенияв него
func withChanel2() {
	var c = make(chan int)
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)

	go func() {
		for {
			select {
			case <-c:
				fmt.Println("\tЗавершение с помощью проброски  канал значения")
				wg.Done()
				return
			default:
				fmt.Println("Heartbeat...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	c <- 69
}

// Заверешение с помощью контекста и функсии cancel
func withContext() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Завершение с помощью контекста")
				wg.Done()
				return
			default:
				fmt.Println("Heartbeat...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
}

// TODO Добавить сюда че нить
