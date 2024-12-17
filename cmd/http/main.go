package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/config"
	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/logger"
	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/storage/postgres"
	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/storage/redis"
)

func main() {
	// init db
	config, err := config.New()
	if err != nil {
		os.Exit(1)
	}
	// set logger
	logger.Set(config.App)
	slog.Info("Start the Application", "app", config.App.Name, "env", config.App.Env)

	// init DB
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

	slog.Info("Successfully connected to the cache server")

	// // Init token service
	// token, err := paseto.New(config.Token)
	// if err != nil {
	// 	os.Exit(1)
	// }

	// r := gin.New()

	// init repository

	// init usecase

	// init handler

	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)

	slog.Info("Starting HTTP Server", "listen_address", listenAddr)
	// slog.Fatal(r.Run(listenAddr))
}
