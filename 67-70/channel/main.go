package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var uploadUrl string

func main() {
	var ch = make(chan string)
	// wg.Add(1)
	go uploadFile(ch)

	fmt.Println(<-ch)
	wg.Wait()
}

func uploadFile(c chan string) {
	fmt.Println("uploading file...")
	time.Sleep(1 * time.Second)
	fmt.Println("File upload done")
	fileUrl := "https://images.pexels.com/photos/35758340/pexels-photo-35758340.jpeg"
	c <- fileUrl
}
