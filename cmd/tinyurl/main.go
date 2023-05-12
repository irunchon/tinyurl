package main

import (
	"database/sql"
	"fmt"
	"github.com/irunchon/tinyurl/internal/pkg/shortening"
	"github.com/irunchon/tinyurl/internal/pkg/storage/inmemory"
	_ "github.com/lib/pq"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "urls_db"
)

func main() {
	storageType := os.Getenv("STORAGE_TYPE")

	switch storageType {
	case "inmemory":
		inmemoryStorageService()
	case "postgres":
		postgresStorageService()
	default:
		fmt.Printf("Unknown storage type\n")
	}
}

func inmemoryStorageService() {
	storage := inmemory.NewInMemoryStorage()
	service := shortening.NewService(storage)

	strings := []string{
		"@@@",
		"!!!",
		"$$$",
		"+++",
		"***",
	}

	for i := range strings {
		storage.SetShortAndLongURLs(service.ShorteningURL(), strings[i])
	}
	fmt.Printf("%v\n", storage)
	val, err := storage.GetShortURLbyLong("!!!")
	fmt.Printf("get for !!!: %v %v\n", val, err)
	val, err = storage.GetShortURLbyLong("***")
	fmt.Printf("get for ***: %v %v\n", val, err)
}

func postgresStorageService() {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
}
