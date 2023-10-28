package domain

// Pet represents the data of a pet.
type Pet struct {
	ID          string      `gorm:"primaryKey;size:300" json:"id"`
	Name        string      `json:"name"`
	Specie      string      `json:"specie"`
	Breed       string      `json:"breed"`
	Color       string      `json:"color"`
	Age         int         `json:"age"`
	OwnerRut    string      `json:"owner_rut" gorm:"size:10"`
	Appointment Appointment `gorm:"foreignKey:PetID"`
}
