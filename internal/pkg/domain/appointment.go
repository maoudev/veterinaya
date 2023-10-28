package domain

import "time"

// Appointment represents information on scheduled appointments at the veterinary clinic.
type Appointment struct {
	ID     string    `gorm:"primaryKey;unique"`
	Date   time.Time `json:"date"`
	PetID  string    `json:"pet_id" gorm:"size:300"`
	VetRut string    `json:"vet_rut" gorm:"size:300"`
	Reason string    `json:"reason"`
	Notes  string    `json:"notes" gorm:"type:text"`
}
