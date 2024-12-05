package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {

	log.Printf("Connected from %s\n", r.RemoteAddr)

	w.Header().Set("Content-type", "text/event-stream")

	w.WriteHeader(http.StatusOK)

	for {

		if f, ok := w.(http.Flusher); ok {
			fmt.Fprintf(w, "%s\n\n", time.Now().Format(time.Stamp))
			f.Flush()
		}
		time.Sleep(1 * time.Second)
	}
}

func savesession() {}

func main() {
	fmt.Printf("Running a server...\n")
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Printf("Can't start the server: %s\n", err)
	}
}
