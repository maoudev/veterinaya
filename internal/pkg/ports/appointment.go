package ports

import "github.com/maoudev/veterinaya/internal/pkg/domain"

type AppointmentRespository interface {
	Create(value interface{}) error
	GetUserAppointments(userRut string) ([]*domain.Appointment, error)
}

type AppointmentService interface {
	Create(appointment *domain.Appointment) error
	GetAppointments(userRut string) ([]*domain.Appointment, error)
}
