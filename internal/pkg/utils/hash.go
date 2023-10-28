package utils

import (
	"strconv"

	"github.com/maoudev/veterinaya/internal/config"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

func TryMatchPasswords(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func HashAndSalt(pass string) string {
	cost, err := strconv.Atoi(config.HASH_COST)
	if err != nil {
		slog.Error(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		slog.Error(err.Error())
	}

	return string(hash)
}
