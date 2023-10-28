package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/maoudev/veterinaya/internal/config"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
)

func CreateToken(claims *domain.Claims) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	strToken, err := jwtToken.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return "", err
	}

	return strToken, nil

}
