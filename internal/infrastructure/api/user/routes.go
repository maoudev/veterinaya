package user

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/infrastructure/repositories/mysql"
	"github.com/maoudev/veterinaya/internal/pkg/service/user"
)

func RegisterRoutes(e *gin.Engine) {
	repo := mysql.NewCLient()
	service := user.NewService(repo)
	handler := newHandler(service)

	e.POST("/api/v1/user", handler.CreateUser)
	e.POST("/api/v1/authenticate", handler.Login)
}
