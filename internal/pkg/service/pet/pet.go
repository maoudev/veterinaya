package pet

import (
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
	"github.com/maoudev/veterinaya/internal/pkg/utils"
)

type petService struct {
	petRepository ports.PetRepository
}

func NewService(repository ports.PetRepository) *petService {
	return &petService{
		petRepository: repository,
	}
}

func (p *petService) Create(pet *domain.Pet) error {
	if !utils.IsRutValid(pet.OwnerRut) {
		return utils.ErrInvalidOwnerRut
	}

	pet.ID = utils.GenerateID()

	return p.petRepository.Create(&pet)
}

func (p *petService) GetPets(userRut string) ([]*domain.Pet, error) {
	pets, err := p.petRepository.GetUserPets(userRut)
	if err != nil {
		return nil, err
	}

	return pets, nil
}
