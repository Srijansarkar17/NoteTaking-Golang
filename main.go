package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Notes struct {
	ID      int
	Title   string
	Content string
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	notes := []Notes{ //making this a list of maps
		{ID: 1,
			Title:   "First Notes",
			Content: "Some description about First Notes"},
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(notes)
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
