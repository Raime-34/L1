package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func sumOfSquares(num []int64) (result int64) {

	var wg sync.WaitGroup
	wg.Add(len(num))
	sum := func(num int64) {
		atomic.AddInt64(&result, num*num) // Использование атомарных операций
		//result = num*num + result
		wg.Done()
	}

	for _, v := range num {
		go sum(v)
	}

	wg.Wait()
	fmt.Println(result)
	return

}

func main() {
	sumOfSquares([]int64{1, 2, 3, 4})
}
