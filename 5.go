package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func spam() {

	// Чтение из консоли
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Было введено нецелочисленное значение")
		return
	}

	var c = make(chan int) // Канал для передачи данных
	var wg sync.WaitGroup
	defer wg.Wait() // Ожидание выполнения горутин
	wg.Add(2)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(n)*time.Second) // Необходима для завершения работы горутин после опредленного периода времени (после n секунд)

	// Анонимная функция для записи в канал
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Запись зарешилась")
				close(c)
				wg.Done()
				return
			default:
				c <- int(time.Now().Unix())
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Анонимная функция для чтения из канала
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Чтение завершилось")
				wg.Done()
				return
			default:
				n, ok := <-c // Проверка закрыт ли канал
				if ok {
					fmt.Println("Принято", n)
				}
			}
		}
	}()

}
