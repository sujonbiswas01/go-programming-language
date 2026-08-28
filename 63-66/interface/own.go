package main

import "fmt"

type Animal1 interface {
	Speek()
}

func make_Sound(a Animal1) {
	a.Speek()
}

type User struct{ name string }

func (u User) Speek() {
	fmt.Println(u.name)
}

func Own() {
	fmt.Println("ownformat")
	u := User{name: "sujon"}
	make_Sound(u)

}
