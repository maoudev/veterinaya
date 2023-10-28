package appointment

import (
	"errors"
	"github.com/google/uuid"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"github.com/maoudev/veterinaya/internal/pkg/ports"
	"github.com/maoudev/veterinaya/internal/pkg/utils"
)

type appointmentService struct {
	repository ports.AppointmentRespository
}

func NewService(repository ports.AppointmentRespository) *appointmentService {
	return &appointmentService{
		repository: repository,
	}
}

func (a *appointmentService) Create(appointment *domain.Appointment) error {
	appointment.ID = uuid.NewString()
	appointment.VetRut = utils.FormatRut(appointment.VetRut)

	_, err := uuid.Parse(appointment.PetID)
	if err != nil {
		return errors.New("error parsing the petid")
	}

	if !utils.IsRutValid(appointment.VetRut) {
		return utils.ErrInvalidVeterinarianRut
	}

	return a.repository.Create(appointment)
}

func (a *appointmentService) GetAppointments(userRut string) ([]*domain.Appointment, error) {
	appointments, err := a.repository.GetUserAppointments(userRut)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}
