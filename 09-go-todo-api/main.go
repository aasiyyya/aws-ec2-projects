package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var todos []Todo
var nextID = 1

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	json.NewDecoder(r.Body).Decode(&t)

	t.ID = nextID
	nextID++

	todos = append(todos, t)

	json.NewEncoder(w).Encode(t)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i := range todos {

		if todos[i].ID == id {

			switch r.Method {

			case "GET":
				json.NewEncoder(w).Encode(todos[i])
				return

			case "PUT":

				var updated Todo
				json.NewDecoder(r.Body).Decode(&updated)

				todos[i].Task = updated.Task
				todos[i].Done = updated.Done

				json.NewEncoder(w).Encode(todos[i])
				return

			case "DELETE":

				todos = append(todos[:i], todos[i+1:]...)

				w.Write([]byte("Todo deleted"))
				return
			}
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case "GET":
			getTodos(w, r)

		case "POST":
			createTodo(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos/", todoHandler)

	fmt.Println("Todo API running on :8085")

	http.ListenAndServe(":8085", nil)
}
