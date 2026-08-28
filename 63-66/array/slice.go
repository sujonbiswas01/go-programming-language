package main

import "fmt"

func Slice() {
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
	// সহজভাবে বললে, Slice (slice) হলো Go-তে এমন একটি dynamic list, যার size প্রয়োজন অনুযায়ী বাড়তে বা কমতে পারে।
	// তাহলে Slice কী?

	// ধরো তুমি জানো না তোমার কতগুলো number লাগবে।
	// Array = নির্দিষ্ট সংখ্যার box 📦
	// Slice = প্রয়োজন অনুযায়ী বড় হওয়া box 📦➡️📦📦
	// 	append() মানে সহজভাবে:

	// Slice-এর শেষে নতুন value যোগ করা।
	// 	len() মানে:

	// Slice-এর মধ্যে বর্তমানে কয়টি element আছে।
	// 	cap() মানে:

	// বর্তমান slice-এর underlying array থেকে সর্বোচ্চ কতটুকু জায়গা ব্যবহার করা সম্ভব, নতুন allocation হওয়ার আগে।
}
