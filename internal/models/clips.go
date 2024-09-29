package models

import (
	"database/sql"
	"errors"
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
    stmt := `SELECT id, title, content, created, expires FROM clips WHERE
        expires > UTC_TIMESTAMP() AND id = ?`

    row := m.DB.QueryRow(stmt, id)

    c := &Clips{}

    err := row.Scan(&c.ID, &c.Title, &c.Content, &c.Created, &c.Expires)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord
        } else {
            return nil, err
        }
    }

    return c, nil
}

func (m *ClipsModel) Latest() ([]*Clips, error) {
    return nil, nil
}

