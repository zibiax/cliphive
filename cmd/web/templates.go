package main

import (
    "github.com/zibiax/cliphive/internal/models"
)

type templateData struct {
    Clip *models.Clip
    Clips []*models.Clip
}
