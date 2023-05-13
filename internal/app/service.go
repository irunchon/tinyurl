package app

import (
	"context"
	"github.com/irunchon/tinyurl/internal/pkg/shortening"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/url"
)

type Service struct {
	pb.UnimplementedShortenURLServer
	repo storage.Storage
}

func New(repo storage.Storage) *Service {
	// TODO: explain UnimplementedShortenURLServer
	return &Service{repo: repo}
}

// GetShortURL == HTTP POST method
func (s Service) GetShortURL(_ context.Context, request *pb.LongURL) (*pb.ShortURL, error) {
	// TODO: check errors and return them to user (if any)

	if !IsUrl(request.LongUrl) {
		// TODO: return error to user if it's not URL
		return nil, status.Errorf(codes.InvalidArgument, "requested URL is not valid")
	}

	var hash string
	// TODO: redo shortening algorithm
	for {
		hash = shortening.GenerateURL()
		if _, err := s.repo.GetLongURLbyShort(hash); err != nil {
			break
		}
	}
	err := s.repo.SetShortAndLongURLs(hash, request.LongUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to add hash and long URL to repository")
	}
	// TODO: check repo behaviour if there are no entities
	return &pb.ShortURL{ShortUrl: hash}, nil
}

// GetLongURL == HTTP GET method
func (s Service) GetLongURL(ctx context.Context, request *pb.Hash) (*pb.LongURL, error) {
	if !IsHashValid(request.Hash) {
		return nil, status.Errorf(codes.InvalidArgument, "requested URL is not valid")
	}

	// TODO: process DB errors
	longURL, err := s.repo.GetLongURLbyShort(request.Hash)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "long URL is not found in repository")
	}
	header := metadata.Pairs("Location", longURL)
	grpc.SendHeader(ctx, header)

	return &pb.LongURL{LongUrl: longURL}, nil
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsHashValid(hash string) bool {
	// TODO: check symbols on validity
	return len(hash) == shortening.ShortURLLength
}
