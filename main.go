package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/infrastructure/api"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	api.RunServer()
}
