package utils

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidOwnerRut        error = errors.New("invalid owner's rut")
	ErrInvalidUserRut         error = errors.New("invalid user's rut")
	ErrInvalidVeterinarianRut error = errors.New("invalid veterinarian's rut")
)

func getRutDigit(value string) (dv string) {
	value = removeUnwantedChars(value)
	rut, _ := strconv.Atoi(value)
	var mlp, acm int
	var counter int = 2
	for rut != 0 {
		mlp = (rut % 10) * counter
		acm += mlp
		rut /= 10
		counter++
		if counter > 7 {
			counter = 2
		}
	}

	digit := 11 - (acm % 11)
	switch 11 - (acm % 11) {
	case 10:
		dv = "K"
	case 11:
		dv = "0"
	default:
		dv = strconv.Itoa(digit)
	}
	return dv
}

func IsRutValid(value string) bool {
	rut := separateDigit(value)
	return getRutDigit(rut[0]) == rut[1]
}

func FormatRut(value string) string {
	value = removeUnwantedChars(value)
	value = strings.ToUpper(value)
	var rutFormatted string
	for i := 0; i < len(value); i++ {
		if i == len(value)-1 {
			rutFormatted += "-" + string(value[i])
			break
		}
		if isCharNumeric(value[i]) {
			rutFormatted += string(value[i])
		}
	}
	return rutFormatted
}

func separateDigit(value string) []string {
	value = removeUnwantedChars(value)
	result := []string{}
	rut := value[:len(value)-1]
	result = append(result, rut)
	digit := value[len(value)-1:]
	result = append(result, digit)

	return result
}

func isCharNumeric(char byte) bool {
	return (int(char) >= 48 && int(char) <= 57)
}

func removeUnwantedChars(value string) string {
	value = strings.TrimSpace(value)
	value = strings.Replace(value, "-", "", -1)
	value = strings.Replace(value, ".", "", -1)
	value = strings.TrimLeft(value, "0")

	return value
}
