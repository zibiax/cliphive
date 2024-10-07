package main

import (
    "net/http"
)

func (app *application) routes() http.Handler {
    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/clip/view", app.cliphiveView)
    mux.HandleFunc("/clip/create", app.cliphiveCreate)

    return app.recoverPanic(app.logRequest(secureHeader(mux)))
}
