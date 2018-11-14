package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "log"
    "github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

    if err != nil {
        log.Fatal(err)
        w.WriteHeader(500)
        return
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(body)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", GetHandler).Methods("GET")
    fmt.Println("Listening on 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
