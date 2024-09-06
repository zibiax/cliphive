package main

import (
    "log"
    "net/http"
)



func main() {
    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", home)
    mux.HandleFunc("/clip/create", cliphiveCreate)
    mux.HandleFunc("/clip/view", cliphiveView)


    log.Println("Starting server on port: 4000")
    println("Yay")

    // Error handling
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
