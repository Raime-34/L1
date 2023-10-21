package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"
)

func opyatRabotat() {

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	n, err := strconv.Atoi(input)

	var list []int        // Объявление контейнера для хранения данных из канала
	var mutex sync.Mutex  // Мютекс понадобится для работы с list
	var wg sync.WaitGroup //
	wg.Add(n + 1)         // Будет запущена еще одна горутина для чтения данных из канала, с дальнейшей передачей их в list

	if err != nil {
		// fmt.Println(err)
		fmt.Println("Было введено нецелочисленное значение")
		return
	}

	emitter := make(chan int, n)                            // Основной канал для постоянного получения данных
	signals := make(chan os.Signal, 1)                      // Дополнительный для работы с Ctrl + C
	ctx, cencel := context.WithCancel(context.Background()) // Понадобится для отслеживание окончания всех горутин

	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		cencel() // Сигнал о необходдимости закончить работу всех горутин
		wg.Wait()
		os.Exit(0)
	}()

	go reader(emitter, &list, &mutex, ctx, &wg) // Запуск рутины для чтения и записи в list

	for i := 0; i < n; i++ {
		go worker(i, &list, &mutex, ctx, &wg) // Запуск n воркеров
	}

	// Цикл, который закидывает данные в канал
	for i := 0; ; i++ {
		emitter <- i
		time.Sleep(50 * time.Microsecond)
	}

}

// Сам воркер
func worker(i int, list *[]int, mutex *sync.Mutex, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %v is dead\n", i+1)
			wg.Done()
			return
		default:
			mutex.Lock()

			if length := len(*list); length > 0 {
				index := rand.Intn(length)
				var n = 0
				*list, n = pull(*list, index)
				fmt.Printf("worker %v получил %v\n", i+1, n)
			}

			mutex.Unlock()
		}
	}
}

func reader(emitter <-chan int, list *[]int, mutex *sync.Mutex, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("reader is dead")
			wg.Done()
			return
		case <-emitter:
			mutex.Lock()
			*list = append(*list, <-emitter)
			mutex.Unlock()
		}
	}
}

// Функция для извлечения(чтения с последующим удалением) элемента
func pull(s []int, i int) ([]int, int) {
	n := s[i]
	s[i] = s[len(s)-1]
	return s[:len(s)-1], n
}
