package main

import "fmt"

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

func main() {
	doSomething()

}
