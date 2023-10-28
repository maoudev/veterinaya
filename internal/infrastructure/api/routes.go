package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/infrastructure/api/appointment"
	"github.com/maoudev/veterinaya/internal/infrastructure/api/pet"
	"github.com/maoudev/veterinaya/internal/infrastructure/api/user"
	"github.com/maoudev/veterinaya/internal/infrastructure/api/veterinarian"
)

func RegisterRoutes(e *gin.Engine) {
	user.RegisterRoutes(e)
	pet.RegisterRoutes(e)
	veterinarian.RegisterRoutes(e)
	appointment.RegisterRoutes(e)
}
