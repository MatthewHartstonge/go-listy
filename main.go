package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Configure Router
	router := NewRouter()
	// Run API Server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}
