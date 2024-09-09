package main

import (
    "log"
    "net/http"
    "strings"
)



func main() {
    mux := http.NewServeMux()



    mux.HandleFunc("/", home)
    mux.HandleFunc("/clip/create", cliphiveCreate)
    mux.HandleFunc("/clip/view", cliphiveView)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))


    log.Println("Starting server on port: 4000")
    println("Yay")

    // Error handling
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}

func neuter(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") {
            http.NotFound(w, r)
            return
        }
        next.ServeHTTP(w, r)
    })
}
