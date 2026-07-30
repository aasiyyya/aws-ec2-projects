package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB

func initDB() {
	var err error

	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User

	json.NewDecoder(r.Body).Decode(&u)

	result, err := db.Exec(
		"INSERT INTO users(name, email) VALUES(?, ?)",
		u.Name,
		u.Email,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	u.ID = int(id)

	json.NewEncoder(w).Encode(u)
}

func main() {
	initDB()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			getUsers(w, r)

		case http.MethodPost:
			createUser(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("SQLite API running on :9090")

	http.ListenAndServe(":9090", nil)
}
