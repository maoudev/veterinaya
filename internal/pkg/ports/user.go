package ports

import "github.com/maoudev/veterinaya/internal/pkg/domain"

type UserRepository interface {
	Create(value interface{}) error
	First(out interface{}, conditions ...interface{}) error
}

type UserService interface {
	Create(user *domain.User) error
	Login(credentials *domain.DefaultCredentials) (string, error)
}
