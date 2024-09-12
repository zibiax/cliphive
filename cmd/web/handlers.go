package main

import (
    "fmt"
    "net/http"
    "strconv"
    "html/template"
)

//Home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    files := []string{
        "./ui/html/pages/base.tmpl",
        "./ui/html/partials/nav.tmpl",
        "./ui/html/pages/home.tmpl",
    }


    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w, err)
    }

}

func (app *application) cliphiveCreate(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    w.Write([]byte("Create a clip..."))
}

func (app *application) cliphiveView(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    
    
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
    fmt.Fprintf(w, "Display a specific clip with ID %d...", id)

}
