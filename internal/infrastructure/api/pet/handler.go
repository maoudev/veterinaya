package pet

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
	"github.com/maoudev/veterinaya/internal/pkg/utils"
)

type petHandler struct {
	petService ports.PetService
}

func newHandler(service ports.PetService) *petHandler {
	return &petHandler{
		petService: service,
	}
}

func (p *petHandler) Add(c *gin.Context) {
	pet := &domain.Pet{}
	if err := c.BindJSON(pet); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	pet.OwnerRut = c.MustGet("Rut").(string)

	err := p.petService.Create(pet)
	if errors.Is(err, utils.ErrInvalidOwnerRut) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (p *petHandler) Get(c *gin.Context) {
	userRut := c.MustGet("Rut").(string)

	pets, err := p.petService.GetPets(userRut)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, pets)
}
