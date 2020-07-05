package users

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

// UserID value object
type UserID struct {
	value string
}

func NewUserId(userId string) (UserID, error) {
	validate := validator.New()
	err := validate.Var(userId, "required,uuid4")
	if err != nil {
		return UserID{}, fmt.Errorf("invalid user id %w", &UserValidationError)
	}
	return UserID{userId}, nil
}

func (uid UserID) IsEqualTo(userId UserID) bool {
	return uid.value == userId.Value()
}

func (uid UserID) Value() string {
	return uid.value
}

// Username value object
type Username struct {
	value string
}

func NewUsername(name string) (Username, error) {
	validate := validator.New()
	err := validate.Var(name, "min=5,max=64")
	if err != nil {
		return Username{}, fmt.Errorf("invalid username: %w", &UserValidationError)
	}
	return Username{name}, nil
}

func (un Username) IsEqualTo(username Username) bool {
	return un.value == username.Value()
}

func (un Username) Value() string {
	return un.value
}
