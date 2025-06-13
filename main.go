package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Gistbox"))
}

func gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific gist with ID %d", id)
	w.Write([]byte(msg))
}

func gistCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to create a new gist"))
}

func gistCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new gist"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /gist/view/{id}", gistView)
	mux.HandleFunc("GET /gist/create", gistCreate)
	mux.HandleFunc("POST /gist/create", gistCreatePost)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
