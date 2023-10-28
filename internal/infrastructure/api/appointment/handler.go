package appointment

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
	"net/http"
)

type appointmentHandler struct {
	appointmentService ports.AppointmentService
}

func newHandler(appointmentService ports.AppointmentService) *appointmentHandler {
	return &appointmentHandler{
		appointmentService: appointmentService,
	}
}

func (a *appointmentHandler) CreateAppointment(c *gin.Context) {
	appointment := &domain.Appointment{}
	if err := c.BindJSON(appointment); err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	if err := a.appointmentService.Create(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusCreated, nil)
}

func (a *appointmentHandler) GetAppointments(c *gin.Context) {
	userRut := c.Query("rut")

	if userRut == "" {
		c.JSON(http.StatusBadRequest, nil)
	}

	appointments, err := a.appointmentService.GetAppointments(userRut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, appointments)
}
