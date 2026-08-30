package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	name string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /createuser", CreateUserHandler)
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
			name: "sujon biswas",
		},
	}
	jsondata, _ := json.Marshal(data)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintln(w, string(jsondata))
}
