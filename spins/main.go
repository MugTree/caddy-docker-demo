package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleSpins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleSpins")
	fmt.Fprintf(w, "Welcome to spins!")
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleNotFound")
	http.NotFound(w, r)
}

func main() {

	r := chi.NewRouter()

	r.Get("/", HandleSpins)
	r.Get("/*", http.NotFound)
	http.ListenAndServe(":8000", r)
}
