package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Rut      string
	Name     string
	LastName string
	Email    string
	Role     string
	jwt.RegisteredClaims
}
