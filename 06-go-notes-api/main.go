package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var notes []Note
var nextID = 1

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	var note Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note.ID = nextID
	nextID++

	notes = append(notes, note)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func main() {

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			getNotes(w, r)

		case http.MethodPost:
			addNote(w, r)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	})

	fmt.Println("Notes API running on :6060")

	http.ListenAndServe(":6060", nil)
}
