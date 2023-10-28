package api

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/maoudev/veterinaya/internal/config"
)

func RunServer() {
	server := gin.Default()

	server.SetTrustedProxies([]string{config.HTTP_ORIGINS})

	server.Use(cors.Middleware(cors.Config{
		Origins:        config.HTTP_ORIGINS,
		Methods:        "GET,PUT,POST,DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		MaxAge:         50 * time.Second,
	}))

	RegisterRoutes(server)

	server.Run(config.API_PORT)
}
