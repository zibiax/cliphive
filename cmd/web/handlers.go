package main

import (
    "fmt"
    "net/http"
    "strconv"
)

//Home handler function
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("Hello from cliphive"))
}

func cliphiveCreate(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Now Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Write([]byte("Create a clip"))
}

func cliphiveView(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    
    
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "Display a specific clip with ID %d...", id)

}
