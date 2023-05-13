package storage

import "errors"

var ErrNotFound = errors.New("not found")

type Storage interface {
	GetLongURLbyShort(string) (string, error)
	SetShortAndLongURLs(string, string) error
}
