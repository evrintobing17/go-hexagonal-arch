package http

import (
	"log/slog"
	"strings"

	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(config *config.HTTP, authHandler AuthHandler) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	// TO DO: CUSTOM VALIDATOR
	// v, ok := binding.Validator.Engine().(*validator.Validate)
	// if ok {

	// }

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
