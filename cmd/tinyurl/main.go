package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/irunchon/tinyurl/internal/pkg/config"

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

func main() {
	serviceConfigParameters := config.InitializeServiceParametersFromEnv()
	var repo storage.Storage

	switch serviceConfigParameters.StorageType {
	case "inmemory":
		repo = inmemory.NewInMemoryStorage()
	case "postgres":
		dbParameters := config.InitializeDBParametersFromEnv()
		db, err := setConnectionToPostgresDB(dbParameters)
		if err != nil {
			log.Fatalf("failed to connect to db: %v", err)
		}
		defer db.Close()
		repo = postgres.NewPostgresStorage(db)
	default:
		log.Fatalf("failed to parse storage type\n")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", serviceConfigParameters.GRPCPort))
	if err != nil {
		log.Fatalf("failed to start listening GRPC port: %v", err)
	}
	grpcServer := grpc.NewServer()

	go func() {
		if err := runGatewayHTTPToGRPC(
			serviceConfigParameters,
			runtime.WithForwardResponseOption(responseHeaderMatcher), // middleware for redirect with HTTP code 301
		); err != nil {
			log.Fatalf("failed to run HTTP to GRPC gateway: %v", err)
		}
	}()

	pb.RegisterShortenURLServer(grpcServer, app.New(repo))
	log.Printf("GRPC server listening at port %s", serviceConfigParameters.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to run GRPC: %v", err)
	}
}

func setConnectionToPostgresDB(dbParameters config.DBParameters) (*sql.DB, error) {
	postgresDBConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbParameters.Host, dbParameters.Port, dbParameters.User, dbParameters.Password, dbParameters.Name)

	db, err := sql.Open("postgres", postgresDBConnection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return db, err
}

func runGatewayHTTPToGRPC(serviceConfigParameters config.ServiceParameters, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()

	mux := runtime.NewServeMux(opts...)

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterShortenURLHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("localhost:%s", serviceConfigParameters.GRPCPort),
		dialOpts,
	)
	if err != nil {
		return err
	}

	log.Printf("HTTP server listening at port %s", serviceConfigParameters.HTTPPort)
	return http.ListenAndServe(fmt.Sprintf(":%s", serviceConfigParameters.HTTPPort), mux)
}

// For redirection if client use HTTP:
func responseHeaderMatcher(_ context.Context, w http.ResponseWriter, grpcResponse proto.Message) error {
	if v, ok := grpcResponse.(*pb.GetLongURLResponse); ok {
		w.Header().Set("Location", v.LongUrl)
		w.WriteHeader(http.StatusMovedPermanently)
	}
	return nil
}
