package utils

import "errors"

var (
	ErrInvalidVeterinarianPhone error = errors.New("invalid veterinarian's phone number")
	ErrInvalidUserPhone         error = errors.New("invalid user's phone number")
)

func IsPhoneNumberValid(number int32) bool {
	return number >= 9 || number <= 9
}
