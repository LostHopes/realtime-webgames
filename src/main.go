package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World from %s", r.RequestURI)
}

func main() {
	fmt.Printf("Running a server...\n")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
