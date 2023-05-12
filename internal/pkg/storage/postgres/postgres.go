package postgres

import (
	"database/sql"
	"errors"
)

type Storage struct {
	db *sql.DB
}

func NewPostgresStorage(postgresDB *sql.DB) *Storage {
	return &Storage{db: postgresDB}
}

func (s *Storage) GetShortURLbyLong(longURL string) (string, error) {
	return "", errors.New("1")
}

func (s *Storage) GetLongURLbyShort(shortURL string) (string, error) {
	return "", errors.New("1")
}

func (s *Storage) SetShortAndLongURLs(shortURL string, longURL string) error {
	_, err := s.db.Exec(`insert into urls (short_url, long_url) values($1, $2)`, shortURL, longURL)
	return err
}
