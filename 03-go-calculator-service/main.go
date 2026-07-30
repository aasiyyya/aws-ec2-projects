package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func add(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	fmt.Fprintf(w, "Result = %d", a+b)
}

func sub(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	fmt.Fprintf(w, "Result = %d", a-b)
}

func mul(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	fmt.Fprintf(w, "Result = %d", a*b)
}

func div(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	if b == 0 {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}

	fmt.Fprintf(w, "Result = %d", a/b)
}

func main() {

	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/mul", mul)
	http.HandleFunc("/div", div)

	fmt.Println("Calculator Service running on :8080")

	http.ListenAndServe(":8080", nil)
}
