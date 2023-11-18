package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	wg := sync.WaitGroup{}

	queue := make(chan bool, 3)

	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()

			queue <- true
			defer func() {
				<-queue
			}()
			req(val, &wg)
		}(val)
	}
	wg.Wait()
}

func req(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}
