package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandlePoker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlePoker")
	fmt.Fprintln(w, "Welcome to poker!")
	fmt.Fprintf(w, "%s", r.URL.RawQuery)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleNotFound")
	http.NotFound(w, r)
}

func main() {

	r := chi.NewRouter()

	r.Get("/", HandlePoker)
	r.Get("/*", http.NotFound)
	http.ListenAndServe(":8001", r)
}
