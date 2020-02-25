package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Should return a list of all books") //TODO implement
	})
	
	fmt.Println("Listening on port 80...")
    http.ListenAndServe(":80", r)
}