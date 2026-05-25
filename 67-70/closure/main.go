package main

import "fmt"

func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor

	}
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println("closure")

	// double := multiplyBy(2)
	// fmt.Println(double(55), "bd")

	increment := counter()

	for i := 0; i < 10; i++ {
		fmt.Println(increment())
	}

}
