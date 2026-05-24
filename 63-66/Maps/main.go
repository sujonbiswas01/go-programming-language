package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
	// metainfo addtionalInfo
}

func main() {
	fmt.Println("maps")

	// myMap := make(map[string]int)
	// myMap["phone"] = 11
	// myMap["email"] = 17
	// fmt.Println(myMap["email"])

	// myMap := map[string]string{
	// 	"name":    "sujon",
	// 	"success": "ok",
	// }
	// fmt.Println(myMap)
	// delete(myMap, "name")
	// fmt.Println(myMap)

	myMap := map[string]user{
		"data": {
			name:  "sujon",
			email: "sujon@gmail.com",
			age:   21,
		},
	}
	fmt.Printf("%+v", myMap["data"])
}
