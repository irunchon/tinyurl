package inmemory

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryStorage() *Storage {
	return &Storage{data: make(map[string]string)}
}

func (s *Storage) Get(str string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, isFound := s.data[str]
	if !isFound {
		return "", ErrNotFound
	}
	return value, nil
}

func (s *Storage) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}
