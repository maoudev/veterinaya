package mysql

import (
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"gorm.io/gorm"
)

type client struct {
	db *gorm.DB
}

func NewCLient() *client {
	return &client{
		db: connect(),
	}
}

func (c *client) Create(value interface{}) error {
	return c.db.Create(value).Error
}

func (c *client) First(dest interface{}, conds ...interface{}) error {
	return c.db.First(dest, conds...).Error
}

func (c *client) GetUserPets(userRut string) ([]*domain.Pet, error) {
	pets := []*domain.Pet{}
	if err := c.db.Find(&pets, "owner_rut = ?", userRut).Error; err != nil {
		return nil, err
	}

	return pets, nil
}

func (c *client) GetUserAppointments(userRut string) ([]*domain.Appointment, error) {
	appointments := []*domain.Appointment{}

	pets := []*domain.Pet{}

	c.db.Find(&pets, "owner_rut = ?", userRut)

	for _, pet := range pets {
		appointment := &domain.Appointment{}
		c.db.First(appointment, "pet_id = ?", pet.ID)

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
