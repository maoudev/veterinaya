package veterinarian

import (
	"errors"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
	"github.com/maoudev/veterinaya/internal/pkg/utils"
)

type veterinarianService struct {
	veterinarianRepository ports.VeterinarianRepository
}

func NewService(repository ports.VeterinarianRepository) *veterinarianService {
	return &veterinarianService{
		veterinarianRepository: repository,
	}
}

func (v *veterinarianService) Create(vet *domain.Veterinarian) error {

	vet.Rut = utils.FormatRut(vet.Rut)

	if !utils.IsRutValid(vet.Rut) {
		return utils.ErrInvalidVeterinarianRut
	}
	_, err := mail.ParseAddress(vet.Email)
	if err != nil {
		return err
	}

	if !utils.IsPhoneNumberValid(vet.Phone) {
		return utils.ErrInvalidUserPhone
	}

	if len(vet.Password) < 8 {
		return errors.New("the password is too short")
	}

	vet.Password = utils.HashAndSalt(vet.Password)

	return v.veterinarianRepository.Create(vet)
}

func (v *veterinarianService) Login(credentials *domain.DefaultCredentials) (string, error) {
	vet := &domain.Veterinarian{}

	if err := v.veterinarianRepository.First(vet, "email = ?", credentials.Email); err != nil {
		return "", err
	}

	if err := utils.TryMatchPasswords(vet.Password, credentials.Password); err != nil {
		return "", err
	}

	claims := createClaims(vet)

	return utils.CreateToken(claims)
}

func createClaims(vet *domain.Veterinarian) *domain.Claims {
	return &domain.Claims{
		Rut:      vet.Rut,
		Name:     vet.Name,
		LastName: vet.LastName,
		Email:    vet.Email,
		Role:     vet.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    vet.Rut,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}
}
