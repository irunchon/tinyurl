include .env
export

.PHONY: all

all:
	go run cmd/tinyurl/main.go