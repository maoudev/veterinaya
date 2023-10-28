package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
)

type userHandler struct {
	userService ports.UserService
}

func newHandler(service ports.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	user := &domain.User{}

	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := u.userService.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (u *userHandler) Login(c *gin.Context) {
	credentials := &domain.DefaultCredentials{}

	if err := c.BindJSON(credentials); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	jwtToken, err := u.userService.Login(credentials)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})
}
