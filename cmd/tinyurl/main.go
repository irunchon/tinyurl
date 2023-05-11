package main

import (
	"fmt"
	"github.com/irunchon/tinyurl/internal/pkg/shortening"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - %s\n", i, shortening.ShorteningURL())
	}
}
