package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string, 3)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "file upload completed"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "file upload"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "file sending.."
	}()

	for range 3 {
		fmt.Println(<-ch)
	}

}
