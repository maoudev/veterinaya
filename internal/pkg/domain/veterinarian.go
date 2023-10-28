package domain

import "time"

// Veterinarian represents a vet and his data.
type Veterinarian struct {
	Rut         string      `gorm:"primaryKey;unique;size:300"`                    // Unique identifier for the veterinarian.
	Name        string      `gorm:"size50" json:"name" binding:"required"`         // First name of the veterinarian.
	LastName    string      `gorm:"size50" json:"last_name" binding:"required"`    // Last name of the veterinarian.
	Email       string      `gorm:"unique" json:"email" binding:"required"`        // Email address of the veterinarian.
	Phone       int32       `gorm:"unique" json:"phone" binding:"required"`        // Phone number of the veterinarian.
	Address     string      `gorm:"size:200" json:"address"`                       // Address of the veterinarian.
	Speciality  string      `gorm:"size:100" json:"speciality" binding:"required"` // Specialization or area of expertise.
	Password    string      `json:"password" binding:"required"`                   // Password with which the veterinarian enters the platform.
	CreatedAt   time.Time   `gorm:"autoCreateTime:true"`                           // Date the vet's account was created.
	Role        string      `gorm:"default:veterinarian" json:"role"`              // Role of the veterinarian account
	Appointment Appointment `gorm:"foreignKey:VetRut"`
}
