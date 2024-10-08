package main

import (
    "fmt"
    "net/http"
    "strconv"
    "errors"

    "github.com/zibiax/cliphive/internal/models"
    "github.com/julienschmidt/httprouter"
)

//Home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {

    clips, err := app.clip.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    data := app.newTemplateData(r)
    data.Clips = clips

    app.render(w, http.StatusOK, "home.tmpl", data)
}
func (app *application) cliphiveCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display the form for creating a new clip"))
}

func (app *application) cliphiveCreatePost(w http.ResponseWriter, r *http.Request) {



    title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
    expires := 7

    id, err := app.clip.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/clip/view/%d", id), http.StatusSeeOther)

}


func (app *application) cliphiveView(w http.ResponseWriter, r *http.Request) {
    params := httprouter.ParamsFromContext(r.Context())
    id, err := strconv.Atoi(params.ByName("id"))
    
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    clip, err := app.clip.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord){
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }
    data := app.newTemplateData(r)
    data.Clip = clip

    app.render(w, http.StatusOK, "view.tmpl", data)
}
