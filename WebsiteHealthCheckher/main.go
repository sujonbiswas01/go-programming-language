package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	url    string
	status string
	Err    error
}

func CheckWebsiteUrl(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
		ch <- Result{url: url, status: "is down", Err: err}
		fmt.Println(<-ch)
	}

	fmt.Println(url, "is Up")
	ch <- Result{url: url, status: "is up and running", Err: nil}
	fmt.Println(<-ch)

	defer res.Body.Close()
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.coms",
	}

	ch := make(chan Result)

	start := time.Now()

	for _, url := range urls {
		go CheckWebsiteUrl(url, ch)
	}
	for range urls {
		result := <-ch
		fmt.Println(result)
	}

	fmt.Println("time taken", time.Since(start))
	fmt.Println("All url checked Successfully")
}
