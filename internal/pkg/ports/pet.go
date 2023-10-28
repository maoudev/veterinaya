package ports

import "github.com/maoudev/veterinaya/internal/pkg/domain"

type PetRepository interface {
	Create(value interface{}) error
	GetUserPets(userRut string) ([]*domain.Pet, error)
}

type PetService interface {
	Create(pet *domain.Pet) error
	GetPets(userRut string) ([]*domain.Pet, error)
}
