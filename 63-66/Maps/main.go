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

	// একটা Map বানালাম → "data" নামে একটা key রাখলাম → সেই key-এর ভিতরে একজন user-এর সম্পূর্ণ information রাখলাম।
	// এই Map-এর key হবে string এবং value হবে user struct।

	myMap := map[string]user{
		"data": {
			name:  "sujon",
			email: "sujon@gmail.com",
			age:   21,
		},
		"d": {
			name:  "rajon",
			email: "rajon@gmail.com",
			age:   17,
		},
	}
	fmt.Printf("%+v", myMap["data"], myMap["d"])
}
