package main

import "fmt"

func main() {
	// variables()
	// datatype()
	mycoffe, inst := Makecofee("black", true, 2)
	Makecofee("cold", false, 5)
	Makecofee("hot", true, 5)
	fmt.Println("i am having ", mycoffe, inst)

}
