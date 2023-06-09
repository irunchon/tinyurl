package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/irunchon/tinyurl/internal/app"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	"github.com/irunchon/tinyurl/internal/pkg/storage/inmemory"
	"github.com/irunchon/tinyurl/internal/pkg/storage/postgres"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// TODO: port, etc. -> env
const (
	//host     = "localhost"
	dbPort   = 5432
	user     = "test"
	password = "test"
	dbname   = "urls_db"
	grpcPort = 50051
	httpPort = 8080
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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	go func() {
		if err := runGatewayHTTPToGRPC(
			httpPort,
			runtime.WithForwardResponseOption(responseHeaderMatcher), // middleware for redirect with HTTP code 301
		); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	pb.RegisterShortenURLServer(grpcServer, app.New(repo))
	log.Printf("GRPC server listening at port %d", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func setConnectionToPostgresDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")

	postgresDBConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, user, password, dbname)

	db, err := sql.Open("postgres", postgresDBConnection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return db, err
}

func runGatewayHTTPToGRPC(serverPort int, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()

	mux := runtime.NewServeMux(opts...)

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterShortenURLHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcPort), dialOpts)
	if err != nil {
		return err
	}

	log.Printf("HTTP server listening at port %d", serverPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", serverPort), mux)
}

// For redirection if client use HTTP:
func responseHeaderMatcher(_ context.Context, w http.ResponseWriter, grpcResponse proto.Message) error {
	if v, ok := grpcResponse.(*pb.GetLongURLResponse); ok {
		w.Header().Set("Location", v.LongUrl)
		w.WriteHeader(http.StatusMovedPermanently)
	}
	return nil
}
