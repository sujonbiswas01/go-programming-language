package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var db *pgx.Conn

func connectDb() {
	var err error
	DATABASE_URL := "postgres://postgres:sujon123@localhost:5432/app"
	db, err = pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected successfully")
}

// var users = []Users{
// 	{
// 		id:    1,
// 		Name:  "sujon biswas",
// 		Age:   20,
// 		Email: "sujon@gmail.com",
// 	},
// 	{
// 		id:    2,
// 		Name:  "rajon biswas",
// 		Age:   15,
// 		Email: "rajon@gmail.com",
// 	},
// }

func main() {
	connectDb()
	defer db.Close(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /health", healHandler)
	mux.HandleFunc("POST /createuser", CreateUserHandler)
	mux.HandleFunc("GET /users", GetuersHandeller)
	// mux.HandleFunc("GET /user/{id}", GetSingleuersHandeller)
	mux.HandleFunc("PUT /user/{id}", UpdateuersHandeller)
	mux.HandleFunc("DELETE /user/{id}", DeleteuersHandeller)

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
	queary := `insert into users (name, age, email) values ($1, $2, $3) returning id`

	err = db.QueryRow(context.Background(), queary, newuser.Name, newuser.Age, newuser.Email).Scan(&newuser.id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "something went wrong, please try again")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newuser)

}

func GetuersHandeller(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(
		context.Background(),
		"SELECT id, name, age, email FROM users",
	)

	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []Users

	for rows.Next() {
		var user Users
		err := rows.Scan(
			&user.id,
			&user.Name,
			&user.Age,
			&user.Email,
		)
		if err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

// func GetSingleuersHandeller(w http.ResponseWriter, r *http.Request) {
// 	idparams := r.PathValue("id")
// 	// fmt.Println(idparams)
// 	// fmt.Fprintln(w, idparams)

// 	// fmt.Printf("the value or id is %v and the type of the id is %T", idparams, idparams)

// 	// for _,user := range users{
// 	// 	if user.id== {
// 	// 	}
// 	// }

// 	id, err := strconv.Atoi(idparams)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintln(w, err)
// 	}

// 	for _, user := range users {
// 		if user.id == id {
// 			json.NewEncoder(w).Encode(user)
// 			return
// 		}
// 	}

// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprintln(w, `not found by id`, id)

// }

func UpdateuersHandeller(w http.ResponseWriter, r *http.Request) {
	idparams := r.PathValue("id")

	id, err := strconv.Atoi(idparams)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}

	var updateduser Users
	err = json.NewDecoder(r.Body).Decode(&updateduser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "invalid user info")
		return
	}

	query := `update users set name=$1, age=$2, email=$3 where id=$4 returning id,name,age,email`

	err = db.QueryRow(context.Background(), query, updateduser.Name, updateduser.Age, updateduser.Email, id).Scan(&updateduser.id, &updateduser.Name, &updateduser.Age, &updateduser.Email)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateduser)

}

func DeleteuersHandeller(w http.ResponseWriter, r *http.Request) {

	idparams := r.PathValue("id")

	id, err := strconv.Atoi(idparams)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}
	var user Users

	query := `
DELETE FROM users
WHERE id = $1
RETURNING id, name, age, email
`
	err = db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.id,
		&user.Name,
		&user.Age,
		&user.Email,
	)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(
		w,
		"Deleted User: %+v\n",
		user,
	)

}
