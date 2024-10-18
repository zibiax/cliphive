package main

import (
    "fmt"
    "net/http"
    "strconv"
    "errors"
    _ "strings"
    _ "unicode/utf8"

    "github.com/zibiax/cliphive/internal/models"
    "github.com/zibiax/cliphive/internal/validator"
    "github.com/julienschmidt/httprouter"
)

type clipCreateForm struct {
    Title string `form:"title"`
    Content string `form:"content"`
    Expires int `form:"expires"`
    validator.Validator `form:"_"`
}

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
    data := app.newTemplateData(r)
    data.Form = clipCreateForm {
        Content: "Roses are red, \nViolets are blue,\nNow it's your turn,\nTo write something new!",
        Expires: 1,
    }

    app.render(w, http.StatusOK, "create.tmpl", data)
}

func (app *application) cliphiveCreatePost(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    var form clipCreateForm
    
    /*
    expires, err := strconv.Atoi(r.PostForm.Get("expires"))
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
    form := clipCreateForm{
        Title: r.PostForm.Get("title"),
        Content: r.PostForm.Get("content"),
        Expires: expires,
    }
    */
    err = app.formDecoder.Decode(&form,r.PostForm)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return

    form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")

    form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")

    form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")

    form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365,), "expires", "This field must equal 1, 7 or 365")

    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
        return
    }



    id, err := app.clip.Insert(form.Title, form.Content, form.Expires)
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
