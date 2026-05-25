package main

// Go এ defer ব্যবহার করা হয় কোনো function call কে পরে execute করার জন্য।
import "fmt"

func defered(result int) {
	fmt.Println("deferret result", result)
}

func example() int {
	result := 10
	defer defered(result)
	fmt.Println("i am comming")
	return result
}

func main() {
	example()
	fmt.Println("defer")
	defer fmt.Println("World")
	fmt.Println("Hello")

}
