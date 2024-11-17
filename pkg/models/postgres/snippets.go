package postgres

import (
	"database/sql"
	"postgre/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string) (int, error) {
	stmt := `INSERT INTO snippets (title, content) VALUES($1, $2) RETURNING id`
	var id int
	err := m.DB.QueryRow(stmt, title, content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
