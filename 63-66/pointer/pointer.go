package main

import "fmt"

// func change(y *int) {
// 	*y = 100
// 	fmt.Printf("%d\n", *y)
// }

func pointer() {
	// fmt.Println("pointer")
	// age := 20

	// var ptr = &age

	// fmt.Printf("%d\n", *ptr)
	// *ptr = 100
	// fmt.Printf("%d\n", *ptr)

	// y := 10
	// change(&y)
	// fmt.Printf("%d\n", y)

	bigArray := [5]int{10, 20, 40, 60, 70}
	modifyWithPointer(&bigArray)
	// modifyWithoutPointer(bigArray)
	// fmt.Println("after without pinter", bigArray)

	fmt.Println("after with pointer", bigArray)

}

// func modifyWithoutPointer(arr [5]int) {
// 	arr[0] = 999
// 	fmt.Println("Inside without pointer", arr)
// }

func modifyWithPointer(arr *[5]int) {
	arr[0] = 777
	fmt.Println("inside the pointer", *arr)
}

// Pointer হলো এমন একটি variable, যেটা অন্য একটি variable-এর memory address (ঠিকানা) ধরে রাখে।
// * এবং & কী?
// Pointer বুঝতে এই দুইটা symbol খুব গুরুত্বপূর্ণ।
// & → Address বের করে
// int age = 20;
// printf("%p", &age);
// &age মানে:
// age-এর memory address কী?
// printf("%p", &age);
// * → Address-এর ভিতরের value বের করে
// int age = 20;

// int *ptr = &age; (C-style)
// fmt.Printf("%d", *ptr) (Go-style)
