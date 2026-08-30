package main

import (
	"fmt"
	"log"
)

func doSomething() {
	defer func() {
		fmt.Println("deferred funtion run")
		r := recover()
		if r != nil {
			fmt.Println("Recoverd form", r)
		}

	}()

	panic("some thing went wronk")
}

func doAnotherThing() {
	defer func() {
		fmt.Println("deferred funtion run")
	}()
	log.Fatal("something very big happend")
}
func main() {
	doSomething()
	doAnotherThing()

}

// 1. log.Fatal কী?

// সহজভাবে:
// log.Fatal = Error দেখাও + সাথে সাথে program বন্ধ করো।

// Program-এর মধ্যে এমন গুরুতর সমস্যা হয়েছে, যেটা normalভাবে handle করা যাচ্ছে না। তাই বর্তমান execution বন্ধ করে panic শুরু করো।"
// Go-তে panic হলেও defer করা function execute হয়।
