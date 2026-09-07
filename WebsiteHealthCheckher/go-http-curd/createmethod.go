package main

import (
	"context"
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
	query := `insert into users (username, age, email) values($1, $2, $3) returning id`

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("user post failed")
	}
	err = db.QueryRow(context.Background(), query, newUser.Name, newUser.Age, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `Could not create user`)
		return
	}
	newUser.ID = len(Datas) + 1
	users := append(Datas, newUser)
	fmt.Println(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}
