package main

import (
    "net/http"
    "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/clip/view", app.cliphiveView)
    mux.HandleFunc("/clip/create", app.cliphiveCreate)

    standard := alice.New(app.recoverPanic, app.logRequest, secureHeader)

    return standard.Then(mux)
}
