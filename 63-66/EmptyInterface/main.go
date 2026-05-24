package main

import "fmt"

func Print(data any) {
	fmt.Println(data)
}

func process(data any) {
	strData, ok := data.(string)
	if !ok {
		fmt.Println("data type not match")
	}
	fmt.Println(strData)
}
func main() {
	fmt.Println("emty interface")
	Print([]int{10, 30})
	Print("sujon")
	process("50")
	process(50)
}
