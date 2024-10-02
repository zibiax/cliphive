package main

import (
    "github.com/zibiax/cliphive/internal/models"
    "html/template"
    "path/filepath"
)

type templateData struct {
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

        files := []string{
            "./ui/html/pages/base.tmpl",
            "./ui/html/partials/nav.tmpl",
            page,
        }
        ts, err := template.ParseFiles(files...)
        if err != nil {
            return nil, err
        }

        cache[name] = ts
    }
    return cache, nil
}

