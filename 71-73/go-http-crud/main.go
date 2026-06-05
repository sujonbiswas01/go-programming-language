package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users = []Users{
	{
		id:    1,
		Name:  "sujon biswas",
		Age:   20,
		Email: "sujon@gmail.com",
	},
	{
		id:    2,
		Name:  "rajon biswas",
		Age:   15,
		Email: "rajon@gmail.com",
	},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /health", healHandler)
	mux.HandleFunc("POST /createuser", CreateUserHandler)
	mux.HandleFunc("GET /users", GetuersHandeller)

	fmt.Println("server is running at http://localhost:5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("server error", err)
		return
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Fprintln(w, "welcome to go server")
}

func healHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is up and healthy")
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "server is up and CreateUserHandler")
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	fmt.Fprintln(w, "only post method is allow")
	// 	return
	// }
	// fmt.Println("user created")
	// fmt.Fprintln(w, "user created")

	var newuser Users

	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "invalid user info")
		return
	}
	newuser.id = len(users) + 1
	fmt.Println(newuser)
	users = append(users, newuser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newuser)

}

func GetuersHandeller(w http.ResponseWriter, r *http.Request) {
	// users, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println("servie")
	// w.Write(users)

	encoder := json.NewEncoder(w)
	encoder.Encode(users)
}
