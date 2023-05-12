package main

import (
	"database/sql"
	"fmt"
	"github.com/irunchon/tinyurl/internal/app"
	"github.com/irunchon/tinyurl/internal/pkg/shortening"
	"github.com/irunchon/tinyurl/internal/pkg/storage"
	"github.com/irunchon/tinyurl/internal/pkg/storage/inmemory"
	"github.com/irunchon/tinyurl/internal/pkg/storage/postgres"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
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

	service := shortening.NewService(repo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterShortenURLServer(grpcServer, &app.Server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//strings := []string{
	//	"@@@",
	//	"!!!",
	//	"$$$",
	//	"+++",
	//	"***",
	//}
	//
	//for i := range strings {
	//	err := repo.SetShortAndLongURLs(service.GenerateURL(), strings[i])
	//	if err != nil {
	//		fmt.Printf("*** %v ***\n", err)
	//	}
	//}
	//fmt.Printf("%v\n", repo)
	//val, err := repo.GetShortURLbyLong("!!!")
	//fmt.Printf("get for !!!: %v %v\n", val, err)
	//val, err = repo.GetShortURLbyLong("***")
	//fmt.Printf("get for ***: %v %v\n", val, err)
	//
	//val, err = repo.GetLongURLbyShort("QLzinpjrTO")
	//fmt.Printf("\nget long for QLzinpjrTO: %v %v\n", val, err)
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
