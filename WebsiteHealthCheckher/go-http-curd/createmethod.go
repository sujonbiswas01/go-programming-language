package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	name  string
	email string
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	fmt.Println("only post method is allow")
	// 	return
	// }

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("user post failed")
	}
	newUser.ID = len(Datas) + 1
	users := append(Datas, newUser)
	fmt.Println(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}
