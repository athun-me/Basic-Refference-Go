package main

import (
	"fmt"
	"sync"
)

func main() {
	var result string
	var wg sync.WaitGroup

	wg.Add(1)

	go test(&result, &wg)

	wg.Wait()

	fmt.Println(result)
}

func test(result *string, wg *sync.WaitGroup) {
	defer wg.Done()
	*result = "hai appssszzosky"
}
