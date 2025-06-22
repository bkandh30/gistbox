package mocks

import (
	"time"

	"gistbox.bhavya.net/internal/models"
)

var mockGist = models.Gist{
	ID:      1,
	Title:   "Mock Gist",
	Content: "This is a mock gist",
	Created: time.Now(),
	Expires: time.Now(),
}

type GistModel struct{}

func (m *GistModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *GistModel) Get(id int) (models.Gist, error) {
	switch id {
	case 1:
		return mockGist, nil
	default:
		return models.Gist{}, models.ErrNoRecord
	}
}

func (m *GistModel) Latest() ([]models.Gist, error) {
	return []models.Gist{mockGist}, nil
}
