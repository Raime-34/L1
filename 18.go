package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i   int            // Само значение котрое мы инкрементируем
	wg  sync.WaitGroup //
	max int            // Максимальное допустимое значение для i
	rw  sync.RWMutex   //
}

func (it *Counter) increment() {
	it.rw.Lock()
	defer it.rw.Unlock()
	it.i++
}

func (it *Counter) print() {
	fmt.Println(it.i)
}

func adsincrement() {

	counter1 := Counter{max: 69}
	counter1.wg.Add(10)
	defer counter1.wg.Wait()

	counter2 := Counter{max: 130}
	counter2.wg.Add(10)
	defer counter2.wg.Wait()

	// Ниже я запускаю 20 горутин для инкрементирования счетчиков

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10 && counter1.i < counter1.max; i++ {
				counter1.increment()
			}
			counter1.wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 60 && counter2.i < counter2.max; i++ {
				counter2.increment()
			}
			counter2.wg.Done()
		}()
	}

	counter1.print()
	counter2.print()
}
