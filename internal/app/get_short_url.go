package app

import (
	"context"
	"errors"
	"net/url"

	"github.com/irunchon/tinyurl/internal/pkg/shortening"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errorAlreadyExists = errors.New("pair URL-hash already exists in repo")

// GetShortURL generates short URL (hash) by long URL (example - HTTP POST method)
// TODO: mock tests for errors in repo
func (s Service) GetShortURL(_ context.Context, request *pb.GetShortURLRequest) (*pb.GetShortURLResponse, error) {
	if !IsUrl(request.LongUrl) {
		return nil, status.Errorf(codes.InvalidArgument, "requested URL is not valid")
	}

	hash, err := s.generateUniqueHashForURL(request.LongUrl)
	if err == errorAlreadyExists {
		return &pb.GetShortURLResponse{ShortUrl: hash}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = s.repo.SetShortAndLongURLs(hash, request.LongUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to add hash and long URL to repository")
	}
	return &pb.GetShortURLResponse{ShortUrl: hash}, nil
}

func IsUrl(stringToCheck string) bool {
	u, err := url.Parse(stringToCheck)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (s Service) generateUniqueHashForURL(longURL string) (string, error) {
	hash := shortening.GenerateHashForURL(longURL)

	for i := 0; i < len(hash)-1; i++ {
		urlToCheck, err := s.repo.GetLongURLbyShort(hash)
		if err == storage.ErrNotFound {
			return hash, nil
		}
		if err != nil {
			return "", errors.New("fail to check hash for duplications")
		}
		if urlToCheck == longURL {
			return hash, errorAlreadyExists
		}
		hash = hashRingShift(hash)
	}
	return "", errors.New("fail to generate hash")
}

func hashRingShift(hash string) string {
	return hash[1:] + hash[:1]
}
