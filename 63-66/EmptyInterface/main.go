package main

import "fmt"

func Print(data any) {
	fmt.Println(data)
}

func process(data any) {
	strData, ok := data.(string)
	fmt.Println(ok)
	if !ok {
		fmt.Println("data type not match")
	}
	fmt.Println(strData)
}

// Go-তে Empty Interface মানে হলো এমন একটা interface যেটা যেকোনো ধরনের value রাখতে পারে।
// কারণ any / interface{}-এর কোনো method requirement নেই।

// তাই Go-র যেকোনো type এটাকে satisfy করতে পারে।
func main() {
	// fmt.Println("emty interface")
	// Print([]int{10, 30})
	// Print("sujon")
	// process("50")
	// process(50)

	var data interface{}
	data = "sujon"
	data = 40
	data = true
	fmt.Println(data)
}
