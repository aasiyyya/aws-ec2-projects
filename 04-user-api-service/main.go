package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User API is running")
}

func users(w http.ResponseWriter, r *http.Request) {
	list := []User{
		{1, "Alice", "Developer"},
		{2, "Bob", "DevOps Engineer"},
		{3, "Charlie", "Cloud Engineer"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/users", users)

	fmt.Println("User API running on :9090")

	http.ListenAndServe(":9090", nil)
}
