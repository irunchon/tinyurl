package inmemory

import (
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	"sync"
)

type Storage struct {
	hashToLongURL map[string]string
	mu            sync.RWMutex
}

func NewInMemoryStorage() *Storage {
	return &Storage{hashToLongURL: make(map[string]string)}
}

func (s *Storage) GetLongURLbyShort(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, isFound := s.hashToLongURL[shortURL]
	if !isFound {
		return "", storage.ErrNotFound
	}
	return value, nil
}

func (s *Storage) SetShortAndLongURLs(shortURL string, longURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hashToLongURL[shortURL] = longURL

	return nil
}
