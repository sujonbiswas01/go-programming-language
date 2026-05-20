package main

import "fmt"

func change(x *int) {
	*x = 400
	fmt.Println("change : ", x)
}

func pointer() {
	fmt.Println("pointer")
	a := 42
	p := &a
	change(p)
	// a = 100
	// fmt.Println("a:", a)
	// fmt.Println("p:", *p)
	// fmt.Println("inside funtion", *p)

	bigArray := [5]int{10, 20, 40, 60, 10}
	// modifyWithoutPointer(bigArray)

	modifyWithPointer(&bigArray)

}

func modifyWithoutPointer(arr [5]int) {
	arr[0] = 1020
	fmt.Println("inside without pointer", arr)
}

func modifyWithPointer(arr *[5]int) {
	arr[0] = 777
	fmt.Println("inside with pointer:", *arr)
}
