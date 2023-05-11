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
	return &Service{storage: s}
}

func (s *Service) ShorteningURL() string {
	for {
		shortURL := make([]rune, shortURLLength)
		for i := 0; i < shortURLLength; i++ {
			shortURL[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
		}
		if _, err := s.storage.GetLongURLbyShort(string(shortURL)); err != nil {
			return string(shortURL)
		}
	}
}
