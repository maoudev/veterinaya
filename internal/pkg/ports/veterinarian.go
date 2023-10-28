package ports

import "github.com/maoudev/veterinaya/internal/pkg/domain"

type VeterinarianRepository interface {
	Create(value interface{}) error
	First(out interface{}, conditions ...interface{}) error
}

type VeterinarianService interface {
	Create(vet *domain.Veterinarian) error
	Login(credentials *domain.DefaultCredentials) (string, error)
}
