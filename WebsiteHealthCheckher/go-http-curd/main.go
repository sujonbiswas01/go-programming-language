package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var Datas = []User{
	{
		ID:    1,
		Name:  "sujon biswas",
		Age:   28,
		Email: "sujon@example.com",
	},
	{
		ID:    2,
		Name:  "rajon biswas",
		Age:   25,
		Email: "rajon@example.com",
	},
	{
		ID:    3,
		Name:  "madob",
		Age:   30,
		Email: "madob@example.com",
	},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /createuser", CreateUserHandler)
	mux.HandleFunc("GET /users", GgetUserhandler)
	// http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"name": "sujonbi",
	}

	jsonData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonData))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]User{
		"su": {
			Name: "sujon biswas",
		},
	}
	jsondata, _ := json.Marshal(data)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintln(w, string(jsondata))
}

func GgetUserhandler(w http.ResponseWriter, r *http.Request) {
	// userss, _ := json.Marshal(datas)
	// w.Write(userss)

	encoder := json.NewEncoder(w)
	encoder.Encode(Datas)
}
