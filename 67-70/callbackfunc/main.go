package main

import "fmt"

func process(callback func()) {
	callback()
}

// Callback হলো এমন একটি function, যেটা অন্য একটি function-এর কাছে argument হিসেবে পাঠানো হয় এবং পরে প্রয়োজন হলে সেটাকে call করা হয়।
func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func main() {
	sayHello := func() {
		fmt.Println("sujon")
	}
	process(sayHello)
	// add := func(n1 int, n2 int) int {
	// 	return n1 + n2
	// }
	// fmt.Println(calculate(10, 90, add))

	// anonymous callback function

	// fmt.Println(calculate(7, 5, func(x, y int) int {
	// 	return x - y
	// }))

}
