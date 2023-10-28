package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtToken *jwt.Token
var err error

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getJWT(c)

		jwtToken, err = parseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !isValidToken(jwtToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		setClaims(c, jwtToken)
	}
}

func IsUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := getRole(jwtToken)
		if role != "user" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func IsVeterinarian() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := getRole(jwtToken)
		if role != "veterinarian" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := getRole(jwtToken)
		if role != "admin" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func getRole(jwtToken *jwt.Token) string {
	claims := jwtToken.Claims.(jwt.MapClaims)
	role := claims["Role"].(string)
	return role
}

func getJWT(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("Authorization")
	jwt := strings.TrimPrefix(authorizationHeader, "Bearer ")
	return jwt
}

func parseToken(strToken string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		return []byte("veterinaya-api"), nil
	})

	if err != nil {
		return nil, errors.New("failed to unmarshal jwt")
	}

	return jwtToken, nil
}

func isValidToken(jwtToken *jwt.Token) bool {
	return jwtToken.Valid
}

func setClaims(c *gin.Context, jwtToken *jwt.Token) {
	claims := jwtToken.Claims.(jwt.MapClaims)

	c.Set("Rut", claims["Rut"])
}
