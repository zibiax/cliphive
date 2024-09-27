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
    stmt := `INSERT INTO clips (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

    result, err := m.DB.Exec(stmt, title, content, expires)
    if err != nil {
        return 0, nil
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, nil
    }
    return int(id), nil
}

func(m *ClipsModel) Get(id int) (*Clips, error) {
    return nil, nil
}

func (m *ClipsModel) Latest() ([]*Clips, error) {
    return nil, nil
}

