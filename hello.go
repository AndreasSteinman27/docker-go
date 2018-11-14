package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    http.Get("https://jsonplaceholder.typicode.com/todos/1")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/metrics", GetHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
