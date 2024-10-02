package main

import (
    "fmt"
    "net/http"
    "strconv"
    "errors"

    "github.com/zibiax/cliphive/internal/models"
)

//Home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    clips, err := app.clip.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }



    app.render(w, http.StatusOK, "home.tmpl", &templateData{Clips: clips})

}

func (app *application) cliphiveCreate(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }
    title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
    expires := 7

    id, err := app.clip.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/clip/view?id=%d", id), http.StatusSeeOther)

}

func (app *application) cliphiveView(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    
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

    app.render(w, http.StatusOK, "view.tmpl", &templateData{Clip: clip})
}
