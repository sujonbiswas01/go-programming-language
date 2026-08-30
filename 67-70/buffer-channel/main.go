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

// ১. Normal Channel
// মানে sender data পাঠালে receiver-এর সাথে synchronization দরকার হয়।

// Buffered Channel

//               Capacity = 3
//                   ↓
//         ┌────┬────┬────┐
// Channel │ 10 │ 20 │ 30 │
//         // └────┴────┴────┘
// অর্থাৎ receiver সাথে সাথে data না নিলেও ৩টা value পর্যন্ত channel-এর ভিতরে রাখা যাবে।
// Buffered Channel = ছোট একটা waiting room/queue, যেখানে receiver নেওয়ার আগে কিছু data অপেক্ষা করতে পারে।
