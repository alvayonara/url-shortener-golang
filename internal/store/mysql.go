package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"url-shortener-golang/models"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(dsn string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLStore{db: db}, nil
}

func (s *MySQLStore) Save(link models.Link) error {
	query := `INSERT INTO links (code, url) VALUES (?, ?)`
	_, err := s.db.Exec(query, link.Code, link.URL)
	return err
}

func (s *MySQLStore) Get(code string) (models.Link, bool) {
	query := `SELECT code, url FROM links WHERE code = ?`
	var link models.Link
	err := s.db.QueryRow(query, code).Scan(&link.Code, &link.URL)
	if err == sql.ErrNoRows {
		return models.Link{}, false
	}
	if err != nil {
		return models.Link{}, false
	}
	return link, true
}
