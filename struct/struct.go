package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
	// metainfo addtionalInfo
}

type addtionalInfo struct {
	phone   int
	address string
}

func structfunc() {
	// fmt.Println("struct")
	// sujon := user{name: "sujon", age: 20, email: "sujon@gmail.com"}
	// sujon.email = "sujon"

	// fmt.Printf("%+v", sujon)

	// var user1 user
	// user1.name = "sujon"
	// user1.email = "sujon@gmail.com"
	// user1.age = 20
	// user1.metainfo.phone = 1788477912
	// user1.metainfo.address = "sylhet"
	// fmt.Printf("%+v", user1)

	newUser := func(name string, email string, age int) user {
		if name == "" {
			fmt.Println("user is required")
			return user{name: "sujon", email: "sujon@gmail.com", age: 20}
		}
		return user{
			name:  name,
			email: email,
			age:   age,
		}
	}
	fmt.Println(newUser("", "sujon@gmail.com", 20))

}
