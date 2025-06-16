package models

import (
	"database/sql"
	"errors"
	"time"
)

type Gist struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type GistModel struct {
	DB *sql.DB
}

func (m *GistModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO gists (title, content, created, expires) 
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *GistModel) Get(id int) (Gist, error) {
	stmt := `SELECT id, title, content, created, expires FROM gists
	WHERE expires > UTC_TIMESTAMP() and id = ?`

	row := m.DB.QueryRow(stmt, id)

	var g Gist

	err := row.Scan(&g.ID, &g.Title, &g.Content, &g.Created, &g.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Gist{}, ErrNoRecord
		} else {
			return Gist{}, err
		}
	}

	return g, nil
}

func (m *GistModel) Latest() ([]Gist, error) {
	return nil, nil
}
