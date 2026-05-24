package main

import "fmt"

func process(callback func()) {
	callback()
}

func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func main() {
	// sayHello := func() {
	// 	fmt.Println("sujon")
	// }
	// process(sayHello)
	add := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println(calculate(10, 90, add))

}
