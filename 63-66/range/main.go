package main

import "fmt"

func main() {
	fmt.Println("range")

	// myMap := map[string]string{
	// 	"name":    "sujon",
	// 	"success": "ok",
	// }

	// for _, value := range myMap {
	// 	fmt.Println(value, "key value")
	// }
	// myArr := [3]string{
	// 	"green",
	// 	"red",
	// 	"black",
	// }

	// for i, value := range myArr {
	// 	fmt.Println(value, i)
	// }

	name := "next level"
	var byteSlice = []byte(name) //কে byte slice-এ convert করা হচ্ছে।
	fmt.Println(byteSlice)

	// Go-তে range ব্যবহার করা হয় কোনো array, slice, string, map, channel ইত্যাদির ভেতরের data একে একে নেওয়ার জন্য।
	// range = collection-এর প্রতিটি item একে একে বের করে আনা।
}
