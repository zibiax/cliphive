package main

import (
    "fmt"
    "net/http"
    "strconv"
    "html/template"
    "log"
)

//Home handler function
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")

    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Errpr", 500)
    }

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
