package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func notesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Note Taking App")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/notes", notesHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	fmt.Println("Server starting on localhost 8080")

}
