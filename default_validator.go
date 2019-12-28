package users

import "errors"

var (
	MinUsernameError = errors.New("username is too sort")
	MaxUsernameError = errors.New("username is too long")
	MinPasswordError = errors.New("too sort password")
	MaxPasswordError = errors.New("too long password")
	InvalidRoleError = errors.New("role can not be empty")
)

func defaultValidator(u User) error {
	minUsernameCharacters := 4
	maxUsernameCharacters := 16
	minPasswordCharacters := 6
	maxPasswordCharacters := 32
	usernameLen := len(u.Username())
	userPassLen := len(u.password)
	if usernameLen < minUsernameCharacters {
		return MinUsernameError
	}
	if usernameLen > maxUsernameCharacters {
		return MaxUsernameError
	}
	if userPassLen < minPasswordCharacters {
		return MinPasswordError
	}
	if userPassLen > maxPasswordCharacters {
		return MaxPasswordError
	}
	if u.GetRole() == "" {
		return InvalidRoleError
	}
	return nil
}