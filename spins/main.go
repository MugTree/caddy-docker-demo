package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleSpins(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to spins!")
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", HandleSpins)
	http.HandleFunc("/*", HandleNotFound)
	http.HandleFunc("/404", HandleNotFound)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
