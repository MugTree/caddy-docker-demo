package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandlePoker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to poker!")
}

func main() {
	http.HandleFunc("/", HandlePoker)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
