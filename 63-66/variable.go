package main

import (
	"fmt"
)

func variables() {
	// var variable_name type = expression
	var name string = "sujon"
	// name = "sujon"
	fmt.Println("string!", name)
	var age int = 20
	fmt.Println("int", age)
	// variablename := value
	fullname := "sujon biswas"
	fmt.Println(fullname)

	var a, b int = 1, 2
	fmt.Println(a, "a", b, "b")

}
