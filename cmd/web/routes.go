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

    dynamic := alice.New(app.sessionManager.LoadAndSave)

    router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
    router.Handler(http.MethodGet, "/clip/view/:id", dynamic.ThenFunc(app.cliphiveView))
    router.Handler(http.MethodGet, "/clip/create", dynamic.ThenFunc(app.cliphiveCreate))
    router.Handler(http.MethodPost, "/clip/create", dynamic.ThenFunc(app.cliphiveCreatePost))


    standard := alice.New(app.recoverPanic, app.logRequest, secureHeader)

    return standard.Then(router)
}
