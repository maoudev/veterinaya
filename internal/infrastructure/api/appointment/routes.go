package appointment

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/infrastructure/api/middlewares"
	"github.com/maoudev/veterinaya/internal/infrastructure/repositories/mysql"
	"github.com/maoudev/veterinaya/internal/pkg/service/appointment"
)

func RegisterRoutes(e *gin.Engine) {
	repo := mysql.NewCLient()
	service := appointment.NewService(repo)
	handler := newHandler(service)

	e.POST("/api/v1/appointment", middlewares.Authenticate(), handler.CreateAppointment)
	e.GET("/api/v1/appointments", middlewares.Authenticate(), handler.GetAppointments)
}
