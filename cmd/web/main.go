package main

import (
    "log"
    "net/http"
    "strings"
    "flag"
    "fmt"
)



func main() {
    // set port address to webserver. Default is 4000
    port := flag.String("port", ":4000", "HTTP network portess")

    flag.Parse()

    if !strings.HasPrefix(*port, ":") {
        *port = fmt.Sprintf(":%s", *port)
    }

    mux := http.NewServeMux()


    mux.HandleFunc("/", home)
    mux.HandleFunc("/clip/create", cliphiveCreate)
    mux.HandleFunc("/clip/view", cliphiveView)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))


    log.Printf("Starting server on port %s", *port)

    // Error handling
    err := http.ListenAndServe(*port, mux)
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
