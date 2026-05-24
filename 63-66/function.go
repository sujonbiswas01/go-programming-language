package main

import "fmt"

func Makecofee(kind string, isSuger bool, s int) (string, int) {
	sa := s
	fmt.Println("cofee", kind, "bool", isSuger, "int", sa)
	coffee := fmt.Sprintf("%s", kind)
	return coffee, sa

}
