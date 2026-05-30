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
	// doSomething()
	doAnotherThing()

}
