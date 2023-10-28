package user

import (
	"errors"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
	"github.com/maoudev/veterinaya/internal/pkg/utils"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewService(repository ports.UserRepository) *userService {
	return &userService{
		userRepository: repository,
	}
}

func (u *userService) Create(user *domain.User) error {

	user.Rut = utils.FormatRut(user.Rut)

	if !utils.IsRutValid(user.Rut) {
		return utils.ErrInvalidUserRut
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	if len(user.Password) < 8 {
		return errors.New("the password is too short")
	}

	user.Password = utils.HashAndSalt(user.Password)

	return u.userRepository.Create(user)
}

func (u *userService) Login(credentials *domain.DefaultCredentials) (string, error) {
	user := &domain.User{}

	if err := u.userRepository.First(user, "email = ?", credentials.Email); err != nil {
		return "", err
	}

	if err := utils.TryMatchPasswords(user.Password, credentials.Password); err != nil {
		return "", err
	}

	claims := createClaims(user)

	return utils.CreateToken(claims)
}

func createClaims(user *domain.User) *domain.Claims {
	return &domain.Claims{
		Rut:      user.Rut,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.Rut,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}
}
