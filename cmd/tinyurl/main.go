package main

import (
	"fmt"
	"github.com/irunchon/tinyurl/internal/pkg/shortening"
	"github.com/irunchon/tinyurl/internal/pkg/storage/inmemory"
	"os"
)

func main() {
	storageType := os.Getenv("STORAGE_TYPE")

	if storageType != "inmemory" {
		fmt.Printf("Wrong storage type!\n")
		return
	}

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
