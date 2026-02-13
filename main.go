package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

// int stores the note-id like notes[1], notes[2] etc.
// Initializes nextID with value 1 Happens once, when the program starts
// we create a mutex here, because without it there will be data corruption like if 10 users are hitting /notes at the same time, then the value of id will be corrupted. without mutex race conditions
var (
	notes  = make(map[int]Note)
	nextID = 1
	mu     sync.Mutex
)

func createNote(w http.ResponseWriter, r *http.Request) {
	var input struct { //structure to recieve json from client
		Title   string
		Content string
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

}
func getNotes(w http.ResponseWriter) {
	mu.Lock()         //now noone can access the notes
	defer mu.Unlock() //when the function ends, the notes unlock automatically

	result := make([]Note, 0, len(notes)) //an empty slice of notes, With enough space to hold all notes

	for _, note := range notes { //loop over every note in the notes map and store them in result
		result = append(result, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// using switch case we are handling both the GET and POST requests together
// if the call is GET, then we just send all the notes back
// if the call is POST, then we call createNote and read data from request body
func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNotes(w)

	case http.MethodPost:
		createNote(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
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
