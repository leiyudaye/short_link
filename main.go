package main

import (
	"short_link/internal/idgen"
	"short_link/internal/shortener"
	"short_link/internal/storage"

	"short_link/api/http/server"
)

func main() {
	service := shortener.NewService(
		&shortener.Options{
			Store:          storage.NewMemoryStore(),
			Generator:      idgen.NewHashGenerator(),
			MaxGenAttempts: 3,
		})

	ser := server.NewServer(service)
	ser.Init()
	ser.Run()
}
