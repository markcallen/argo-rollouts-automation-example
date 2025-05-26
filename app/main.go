package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/good", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Everything is fine!\n")
    })

    http.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Something went wrong!", http.StatusInternalServerError)
    })

    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
