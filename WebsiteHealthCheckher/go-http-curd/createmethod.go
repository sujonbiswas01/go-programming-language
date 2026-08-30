package main

import (
	"fmt"
	"net/http"
)

type Users struct {
	name  string
	email string
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("only post method is allow")
		return
	}

}
