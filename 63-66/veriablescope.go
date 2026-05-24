package main

import "fmt"

func veriableScope() {
	sugar := 2
	makeCofee := func() {
		sugar := 3
		cofee := "sugar"
		fmt.Printf("making %s with %d spon of sur", cofee, sugar)
	}
	makeCofee()
	x := 5
	fmt.Printf("Inside Scope:", x, sugar)
}
