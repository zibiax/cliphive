package main

import (
    "log"
    "net/http"
    // "fmt"
)

//Home handler function
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from cliphive"))
}


func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)


    log.Println("Starting server on port: 4000")

    // Error handling
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
