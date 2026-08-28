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

// Function-কে নির্দিষ্ট Dog, Cat, Bkash না দিয়ে interface দিলে ভবিষ্যতে একই interface follow করা নতুন type সহজেই ব্যবহার করা যায়।

type Animal interface {
	speek()
}

func makeSound(d Animal) {
	d.speek()
}

type Payment interface {
	Pay()
}

type Bkash struct {
	apikey string
}

func (b Bkash) Pay() {
	fmt.Println("Paid with Bkash")
}
func CompletePayment(p Payment) {
	p.Pay()
}

type PaymentService struct {
}

func main() {
	// fmt.Println("interface")
	// dexter := Animal{}
	// makeSound(dexter)
	// fmt.Println(dexter, "dexer")

	// b := Bkash{}
	// CompletePayment(b)
	Own()
}
