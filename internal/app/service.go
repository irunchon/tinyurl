package app

import (
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
)

type Service struct {
	pb.UnimplementedShortenURLServer
	repo storage.Storage
}

func New(repo storage.Storage) *Service {
	return &Service{repo: repo}
}
