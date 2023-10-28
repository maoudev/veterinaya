package veterinarian

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
)

type veterinarianHandler struct {
	veterinarianService ports.VeterinarianService
}

func newHandler(service ports.VeterinarianService) *veterinarianHandler {
	return &veterinarianHandler{
		veterinarianService: service,
	}
}

func (v *veterinarianHandler) CreateVeterinarian(c *gin.Context) {
	vet := &domain.Veterinarian{}

	if err := c.BindJSON(vet); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := v.veterinarianService.Create(vet); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (v *veterinarianHandler) Login(c *gin.Context) {
	credentials := &domain.DefaultCredentials{}

	if err := c.BindJSON(credentials); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	jwtToken, err := v.veterinarianService.Login(credentials)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})
}
