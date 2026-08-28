package main

import "fmt"

// func sumTwo(a ...int) int {
// 	total := 0

// 	for _, value := range a {
// 		total += value
// 	}
// 	return total
// }

// func greet(prefix string, nps ...string) {
// 	fmt.Println(prefix, nps)
// }

// এমন একটি function, যেটা যত খুশি সংখ্যক argument নিতে পারে।

func add(numbers ...int) int {
	var total = 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func main() {
	// fmt.Println("variaticfun")
	// sum := sumTwo(10, 20, 50, 90, 690)
	// fmt.Println(sum)
	// greet("welcome", "sujon", "jamal")

	fmt.Println(add(10, 50, 60, 870))
}
