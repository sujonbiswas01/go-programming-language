package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Url    string
	Status string
	Err    error
}

func checkWebsiteUrl(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil {
		ch <- Result{Url: url, Status: "down", Err: err}
		return
	}

	ch <- Result{Url: url, Status: "up", Err: nil}

	defer res.Body.Close()
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://sujonbiswasdev.vercel.app",
	}

	ch := make(chan Result)
	start := time.Now()
	for _, url := range urls {
		go checkWebsiteUrl(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Println("All url time", time.Since(start))
}
