package app

import (
	"context"
	"regexp"

	"github.com/irunchon/tinyurl/internal/pkg/shortening"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var regexValidHashChars = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

// GetLongURL searches long URL by hash in storage (example - HTTP GET method)
func (s Service) GetLongURL(ctx context.Context, request *pb.GetLongURLRequest) (*pb.GetLongURLResponse, error) {
	if !isHashValid(request.Hash) {
		return nil, status.Errorf(codes.InvalidArgument, "requested URL is not valid")
	}

	longURL, err := s.repo.GetLongURLbyShort(request.Hash)
	if err == storage.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "long URL is not found in repository")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to get long URL from repository")
	}

	return &pb.GetLongURLResponse{LongUrl: longURL}, nil
}

func isHashValid(hash string) bool {
	return len(hash) == shortening.ShortURLLength && regexValidHashChars.MatchString(hash)
}
