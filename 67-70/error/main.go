package main

import (
	"errors"
	"fmt"
)

var err error

func main() {
	a := DoSomething()
	fmt.Println(a)

	data, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func DoSomething() string {
	fmt.Println("something ")
	return "jfsldkjfsdfdsf"
}
