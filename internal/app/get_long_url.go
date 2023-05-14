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
)

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

func IsHashValid(hash string) bool {
	// TODO: check symbols on validity
	return len(hash) == shortening.ShortURLLength
}
