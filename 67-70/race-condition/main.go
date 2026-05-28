package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	for range 1000 {
		wg.Go(increment)

	}
	wg.Wait()
	fmt.Println("counter value is", counter)
}

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter = counter + 1

}
