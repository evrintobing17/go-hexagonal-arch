package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/config"
	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/storage/redis"

	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/storage/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	config, err := config.New()
	if err != nil {
		os.Exit(1)
	}
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		os.Exit(1)
	}
	db.Close()

	// init redis
	cache, err := redis.New(ctx, config.Redis)
	if err != nil {
		os.Exit(1)
	}
	defer cache.Close()

	// slog.Info("Successfully connected to the cache server")

	// // Init token service
	// token, err := paseto.New(config.Token)
	// if err != nil {
	// 	os.Exit(1)
	// }

	r := gin.New()

	// init repository

	// init usecase

	// init handler

	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)

	log.Printf("Starting server on port %v", config.HTTP.Port)
	log.Fatal(r.Run(listenAddr))

}
