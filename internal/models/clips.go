package models

import (
    "database/sql"
    "time"
)

type Clips struct {
    ID int
    Title string
    Content string
    Created time.Time
    Expires time.Time
}

type ClipsModel struct {
    DB *sql.DB
}

func (m *ClipsModel) Insert(title string, content string, expires int) (int, error) {
    return 0, nil
}

func(m *ClipsModel) Get(id int) (*Clips, error) {
    return nil, nil
}

func (m *ClipsModel) Latest() ([]*Clips, error) {
    return nil, nil
}

