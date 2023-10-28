package domain

// DefaultCredentials represents the email/password combination.
type DefaultCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
