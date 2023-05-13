package postgres

import (
	"database/sql"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
)

type Storage struct {
	db *sql.DB
}

func NewPostgresStorage(postgresDB *sql.DB) *Storage {
	return &Storage{db: postgresDB}
}

func (s *Storage) GetLongURLbyShort(shortURL string) (string, error) {
	row := s.db.QueryRow(`SELECT long_url FROM urls WHERE short_url=$1`, shortURL)
	if row.Err() != nil {
		return "", row.Err()
	}

	var longURL string
	err := row.Scan(&longURL)
	if err == sql.ErrNoRows {
		return "", storage.ErrNotFound
	}

	return longURL, err
}

func (s *Storage) SetShortAndLongURLs(shortURL string, longURL string) error {
	_, err := s.db.Exec(`insert into urls (short_url, long_url) values($1, $2)`, shortURL, longURL)
	return err
}
