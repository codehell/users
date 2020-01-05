package users

import "errors"

var (
	ErrMinUsername = errors.New("username is too sort")
	ErrMaxUsername = errors.New("username is too long")
	ErrMinPassword = errors.New("too sort password")
	ErrMaxPassword = errors.New("too long password")
	ErrInvalidRole = errors.New("role can not be empty")
)

func DefaultValidator(u *User) error {
	minUsernameCharacters := 4
	maxUsernameCharacters := 16
	minPasswordCharacters := 6
	maxPasswordCharacters := 32
	usernameLen := len(u.Username)
	userPassLen := len(u.Password)
	if usernameLen < minUsernameCharacters {
		return ErrMinUsername
	}
	if usernameLen > maxUsernameCharacters {
		return ErrMaxUsername
	}
	if userPassLen < minPasswordCharacters {
		return ErrMinPassword
	}
	if userPassLen > maxPasswordCharacters {
		return ErrMaxPassword
	}
	if u.Role == "" {
		return ErrInvalidRole
	}
	return nil
}
