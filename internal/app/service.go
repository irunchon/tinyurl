package app

import (
	"context"
	"net/url"

	"github.com/irunchon/tinyurl/internal/pkg/shortening"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedShortenURLServer
	repo storage.Storage
}

func New(repo storage.Storage) *Service {
	return &Service{repo: repo}
}

// GetShortURL generates short URL (hash) by long URL (example - HTTP POST method)
func (s Service) GetShortURL(_ context.Context, request *pb.LongURL) (*pb.ShortURL, error) {
	if !IsUrl(request.LongUrl) {
		return nil, status.Errorf(codes.InvalidArgument, "requested URL is not valid")
	}

	hash := shortening.GenerateHashForURL(request.LongUrl)
	isHashOK := false
	for i := 0; i < len(hash)-1; i++ {
		urlToCheck, err := s.repo.GetLongURLbyShort(hash)
		if err == storage.ErrNotFound {
			isHashOK = true
			break
		}
		if err != nil {
			return nil, status.Errorf(codes.Internal, "fail to check hash for duplications")
		}
		if urlToCheck == request.LongUrl {
			return &pb.ShortURL{ShortUrl: hash}, nil
		}
		hash = hashRingShift(hash)
	}
	if !isHashOK {
		return nil, status.Errorf(codes.Internal, "fail to generate hash")
	}

	err := s.repo.SetShortAndLongURLs(hash, request.LongUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to add hash and long URL to repository")
	}
	return &pb.ShortURL{ShortUrl: hash}, nil
}

// GetLongURL searches long URL by hash in storage (example - HTTP GET method)
func (s Service) GetLongURL(ctx context.Context, request *pb.Hash) (*pb.LongURL, error) {
	if !IsHashValid(request.Hash) {
		return nil, status.Errorf(codes.InvalidArgument, "requested URL is not valid")
	}

	longURL, err := s.repo.GetLongURLbyShort(request.Hash)
	if err == storage.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "long URL is not found in repository")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to get long URL from repository")
	}
	header := metadata.Pairs("Location", longURL)
	err = grpc.SendHeader(ctx, header)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to send redirect")
	}

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

func hashRingShift(hash string) string {
	return hash[1:] + hash[:1]
}
