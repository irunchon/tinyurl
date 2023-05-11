package inmemory

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Storage struct {
	keyShortURL map[string]string
	keyLongURL  map[string]string
	mu          sync.RWMutex
}

func NewInMemoryStorage() *Storage {
	return &Storage{
		keyShortURL: make(map[string]string),
		keyLongURL:  make(map[string]string),
	}
}

func (s *Storage) GetShortURLbyLong(longURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, isFound := s.keyLongURL[longURL]
	if !isFound {
		return "", ErrNotFound
	}
	return value, nil
}

func (s *Storage) GetLongURLbyShort(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, isFound := s.keyShortURL[shortURL]
	if !isFound {
		return "", ErrNotFound
	}
	return value, nil
}

func (s *Storage) SetShortAndLongURLs(shortURL string, longURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keyLongURL[longURL] = shortURL
	s.keyShortURL[shortURL] = longURL
}
