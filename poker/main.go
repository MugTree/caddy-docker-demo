package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandlePoker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to poker!")
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", HandlePoker)
	http.HandleFunc("/*", HandleNotFound)
	http.HandleFunc("/404", HandleNotFound)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
