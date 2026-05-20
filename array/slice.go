package main

import "fmt"

func slice() {
	fmt.Println("slice")
	var orders = [6]int{10, 20, 30, 40, 50, 60}
	slice1 := orders[0:4]

	slice1 = append(slice1, 400)
	slice1 = append(slice1, 500)
	slice1 = append(slice1, 600)

	fmt.Println(slice1)
	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Println(orders)
	fmt.Println("the length of the slice is ", len(slice1))
	fmt.Println("the cap of the slice is", cap(slice1))
}
