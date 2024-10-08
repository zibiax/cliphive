package main

import (
    "net/http"
    "github.com/justinas/alice"
    "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
    router := httprouter.New()

    router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.notFound(w)
    })

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

    router.HandlerFunc(http.MethodGet, "/", app.home)
    router.HandlerFunc(http.MethodGet, "/clip/view/:id", app.cliphiveView)
    router.HandlerFunc(http.MethodGet, "/clip/create", app.cliphiveCreate)
    router.HandlerFunc(http.MethodPost, "/clip/create", app.cliphiveCreatePost)


    standard := alice.New(app.recoverPanic, app.logRequest, secureHeader)

    return standard.Then(router)
}
