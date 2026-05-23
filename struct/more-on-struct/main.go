package main

import "fmt"

// type user struct {
// 	name       string
// 	age        int
// 	isLoggedIn bool
// 	greet      func()
// }

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	fmt.Println("more func on struct")
	// user1 := user{

	// 	name:       "sujon",
	// 	age:        20,
	// 	isLoggedIn: false,
	// 	greet: func() {
	// 		fmt.Println("hello there")
	// 	},
	// }

	// user1.greet = func() {
	// 	fmt.Println("hello", user1.name)
	// }

	// user1.greet()

	user1 := user{
		name:       "sujon",
		age:        20,
		isLoggedIn: false,
	}
	// user2 := user{
	// 	name:       "rajon",
	// 	age:        15,
	// 	isLoggedIn: false,
	// }
	user1.greet()
	user1.login()
	fmt.Printf("%+v", user1)
}

func (u *user) login() {
	fmt.Println(`login called`)
	(*u).isLoggedIn = true
}

func (u user) greet() {
	fmt.Println("hello", u.name)
}
