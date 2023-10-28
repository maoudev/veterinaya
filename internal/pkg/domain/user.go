package domain

// User represents a common user.
type User struct {
	Rut      string `gorm:"primaryKey;unique;size:10" json:"rut" binding:"required"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Phone    int32  `json:"phone" gorm:"unique" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `gorm:"default:user" json:"role"`
	Pet      Pet    `gorm:"foreignKey:OwnerRut"`
}
