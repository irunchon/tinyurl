package app

import (
	"context"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: rename to service ???
type Server struct {
	pb.UnimplementedShortenURLServer
	repo storage.Storage
}

func New(repo storage.Storage) *Server {
	// TODO: explain UnimplementedShortenURLServer
	return &Server{repo: repo}
}

func (s Server) GetShortURLbyLong(_ context.Context, url *pb.LongURL) (*pb.ShortURL, error) {
	// TODO: check if URL is real
	// TODO: retrun error to user if it's not URL
	// TODO: generate short URL
	// TODO: check if there are dublications in the repo
	// TODO: if yes - generate new Short URL until it's unique
	// TODO: check repo behaviour if there are no entities
	// TODO: add pair Long URL + Short URL to repo
	// TODO: send short URL to client

	// TODO: check errors
	// TODO: retrun error to user if any (?) problems

	//err := s.repo.SetShortAndLongURLs("", strings[i])

	return nil, status.Errorf(codes.Unimplemented, "method GetShortURLbyLong not implemented")
}
func (s Server) GetLongURLbyShort(_ context.Context, url *pb.ShortURL) (*pb.LongURL, error) {
	// TODO: check if URL is real
	// TODO: check if short URL is in repo
	// TODO: if no return error to client
	// TODO: if ok get long url from repo
	// TODO: return long url to client
	// TODO: if HTTP add header with code 302 and location (long URL)

	// TODO: check errors
	// TODO: retrun error to user if any (?) problems
	return nil, status.Errorf(codes.Unimplemented, "method GetLongURLbyShort not implemented")
}
