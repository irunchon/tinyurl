package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/irunchon/tinyurl/internal/app"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	"github.com/irunchon/tinyurl/internal/pkg/storage/inmemory"
	"github.com/irunchon/tinyurl/internal/pkg/storage/postgres"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"net/http"
	"os"
)

// TODO: port -> env
const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "urls_db"
	grpcPort = "50051"
	httpPort = "8080"
)

// TODO: error processing
func main() {
	storageType := os.Getenv("STORAGE_TYPE")
	var repo storage.Storage

	switch storageType {
	case "inmemory":
		repo = inmemory.NewInMemoryStorage()
	case "postgres":
		db, err := setConnectionToPostgresDB()
		if err != nil {
			panic(err)
		}
		defer db.Close()
		repo = postgres.NewPostgresStorage(db)
	default:
		fmt.Printf("Unknown storage type\n")
		return
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	go func() {
		if err := runGatewayHTTPToGRPC(fmt.Sprintf(
			"localhost:%s", httpPort),
			runtime.WithForwardResponseOption(responseHeaderMatcher), // middleware for redirect with HTTP code 302
		); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	pb.RegisterShortenURLServer(grpcServer, app.New(repo))
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func setConnectionToPostgresDB() (*sql.DB, error) {
	postgresDBConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", postgresDBConnection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return db, err
}

func runGatewayHTTPToGRPC(httpServerAddress string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()

	mux := runtime.NewServeMux(opts...)

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterShortenURLHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", grpcPort), dialOpts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(httpServerAddress, mux)
}

// For redirection if client use HTTP:
func responseHeaderMatcher(_ context.Context, w http.ResponseWriter, _ proto.Message) error {
	headers := w.Header()
	if location, ok := headers["Grpc-Metadata-Location"]; ok {
		w.Header().Set("Location", location[0])    // URL is stored in location[0])
		w.WriteHeader(http.StatusMovedPermanently) // HTTP code 301
	}
	return nil
}
