package pet

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/infrastructure/api/middlewares"
	"github.com/maoudev/veterinaya/internal/infrastructure/repositories/mysql"
	"github.com/maoudev/veterinaya/internal/pkg/service/pet"
)

func RegisterRoutes(e *gin.Engine) {
	repo := mysql.NewCLient()
	service := pet.NewService(repo)
	handler := newHandler(service)

	e.POST("/api/v1/pets", middlewares.Authenticate(), handler.Add)
	e.GET("/api/v1/pets", middlewares.Authenticate(), handler.Get)
}
