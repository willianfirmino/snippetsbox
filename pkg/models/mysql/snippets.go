package mysql

import (
	"database/sql"
	"snippetbox/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil

}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil

}
