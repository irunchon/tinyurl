package shortening

import (
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	"math/rand"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const shortURLLength = 10

type Service struct {
	storage storage.Storage
}

func NewService(s storage.Storage) *Service {
	return &Service{s}
}

func (s *Service) ShorteningURL() string {
	for {
		shortURL := make([]rune, shortURLLength)
		for i := 0; i < shortURLLength; i++ {
			shortURL[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
		}
		if _, errNotFound := s.storage.GetLongURLbyShort(string(shortURL)); errNotFound != nil {
			return string(shortURL)
		}
	}
}
