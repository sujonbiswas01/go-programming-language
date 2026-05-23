package main

import "fmt"

type Dog struct{}
type Cat struct{}

func (d Dog) speek() {
	fmt.Println("Woof")
}

func (d Cat) speek() {
	fmt.Println("Cat")
}

type Animal interface {
	speek()
}

func makeSound(d Animal) {
	d.speek()
}

type Payment interface {
	Pay()
}

type Bkash struct{}

func (b Bkash) Pay() {
	fmt.Println("Paid with Bkash")
}
func CompletePayment(p Payment) {
	p.Pay()
}
func main() {
	// fmt.Println("interface")
	// dexter := Animal{}
	// makeSound(dexter)
	// fmt.Println(dexter, "dexer")

	b := Bkash{}
	CompletePayment(b)
}
