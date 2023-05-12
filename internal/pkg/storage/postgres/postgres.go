package postgres

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func NewPostgresStorage(postgresDB *sql.DB) *Storage {
	return &Storage{db: postgresDB}
}

func (s *Storage) GetShortURLbyLong(longURL string) (string, error) {
	row := s.db.QueryRow(`SELECT short_url FROM urls WHERE long_url=$1`, longURL)
	if row.Err() != nil {
		return "", row.Err()
	}

	var shortURL string
	err := row.Scan(&shortURL)

	return shortURL, err
}

func (s *Storage) GetLongURLbyShort(shortURL string) (string, error) {
	row := s.db.QueryRow(`SELECT long_url FROM urls WHERE short_url=$1`, shortURL)
	if row.Err() != nil {
		return "", row.Err()
	}

	var longURL string
	err := row.Scan(&longURL)

	return longURL, err
}

func (s *Storage) SetShortAndLongURLs(shortURL string, longURL string) error {
	_, err := s.db.Exec(`insert into urls (short_url, long_url) values($1, $2)`, shortURL, longURL)
	return err
}
