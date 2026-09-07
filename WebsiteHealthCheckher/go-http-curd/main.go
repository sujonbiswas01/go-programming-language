package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
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
var db *pgx.Conn

func connectDb() {
	var err error
	urlExample := "postgres://postgres:sujon123@localhost:5432/postgres"
	db, err = pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(db)
	fmt.Println("database connection successfully")

}

func main() {
	connectDb()
	defer db.Close(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /createuser", CreateUserHandler)
	mux.HandleFunc("GET /users", GgetUserhandler)
	mux.HandleFunc("GET /users/{id}", GetSingleuserHandler)
	mux.HandleFunc("PUT /users/{id}", UpdateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", DeleteUserHandler)
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

	// encoder := json.NewEncoder(w)
	// encoder.Encode(Datas)

	rows, err := db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Database-এর users রাখার জন্য Slice
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email)
		if err != nil {
			http.Error(w, "Failed to read user data", http.StatusInternalServerError)
			return
		}
		// User-কে users slice-এ যোগ করা
		users = append(users, user)
		// Response JSON format করা
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func GetSingleuserHandler(w http.ResponseWriter, r *http.Request) {
	idParams := r.PathValue("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("get single user failed")
		return
	}
	for _, user := range Datas {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	// URL থেকে id নেওয়া
	idParam := r.PathValue("id")

	// String id কে int-এ convert করা
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// নতুন তথ্য রাখার জন্য একটি variable
	var updateUser User

	// Request body থেকে JSON data নেওয়া
	err = json.NewDecoder(r.Body).Decode(&updateUser)

	if err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	// সব user-এর মধ্যে খোঁজা
	for index, user := range Datas {

		// URL-এর ID এবং User-এর ID একই কিনা
		if user.ID == id {

			// পুরোনো ID রেখে দেওয়া
			updateUser.ID = id

			// পুরোনো user-এর জায়গায় নতুন user রাখা
			Datas[index] = updateUser

			// Success response পাঠানো
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateUser)

			return
		}
	}

	// কোনো user পাওয়া না গেলে
	http.Error(w, "User not found", http.StatusNotFound)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	// URL থেকে ID নেওয়া
	idParam := r.PathValue("id")

	// String থেকে int-এ convert করা
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Datas-এর মধ্যে User খোঁজা
	for index, user := range Datas {

		if user.ID == id {

			// ওই User-কে Slice থেকে remove করা
			Datas = append(
				Datas[:index],
				Datas[index+1:]...,
			)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "User deleted successfully")

			return
		}
	}
	// User পাওয়া না গেলে
	http.Error(w, "User not found", http.StatusNotFound)
}
