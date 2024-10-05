package main

import (
    "github.com/zibiax/cliphive/internal/models"
    "html/template"
    "path/filepath"
)

type templateData struct {
    CurrentYear int
    Clip *models.Clip
    Clips []*models.Clip
}

func newTemplateCache() (map[string]*template.Template, error) {
    cache := map[string]*template.Template{}

    pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        name := filepath.Base(page)

        ts, err := template.ParseFiles("./ui/html/pages/base.tmpl")
        if err != nil {
            return nil, err
        }
        ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
        if err != nil {
            return nil, err
        }
        ts, err = ts.ParseFiles(page)
        if err != nil {
            return nil, err
        }

        cache[name] = ts
    }
    return cache, nil
}

