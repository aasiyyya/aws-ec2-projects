package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./files"))

	http.Handle("/", fs)

	fmt.Println("File Server running on :7070")

	http.ListenAndServe(":7070", nil)
}
