package main

import (
    "log"
    "net/http"
    // "fmt"
)

//Home handler function
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("Hello from cliphive"))
}

func create(w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("Creating a clip"))
}

func view(w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("view a clip"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/clip/create", create)
    mux.HandleFunc("/clip/view", view)


    log.Println("Starting server on port: 4000")

    // Error handling
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
