package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleSpins(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HandleSpins")
	fmt.Fprintln(w, "Welcome to spins!")
	fmt.Fprintf(w, "%s", r.URL.RawQuery)
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
