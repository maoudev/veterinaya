package veterinarian

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/infrastructure/api/middlewares"
	"github.com/maoudev/veterinaya/internal/infrastructure/repositories/mysql"
	"github.com/maoudev/veterinaya/internal/pkg/service/veterinarian"
)

func RegisterRoutes(e *gin.Engine) {
	repo := mysql.NewCLient()
	service := veterinarian.NewService(repo)
	handler := newHandler(service)

	e.POST("/api/v1/vet", middlewares.Authenticate(), middlewares.IsAdmin(), handler.CreateVeterinarian)
	e.POST("/api/v1/vet/authenticate", handler.Login)
}
