package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleSpins(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to spins!")
}

func main() {
	http.HandleFunc("/", HandleSpins)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
