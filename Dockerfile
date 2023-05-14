FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o app cmd/tinyurl/main.go

ENV STORAGE_TYPE=inmemory

EXPOSE 8080
EXPOSE 50051

ENTRYPOINT ["./app"]