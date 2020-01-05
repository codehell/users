package users

import (
	"github.com/alexedwards/argon2id"
	"time"
)

type User struct {
	ID        interface{} `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Client interface {
	Create(*User) error
	Close() error
	GetUserByEmail(string) (User, error)
	GetAll() ([]User, error)
	DeleteAll() error
	SetValidator(Validator)
}

type Validator func(user *User) error

func CheckPassword(u User, password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, u.Password)
	if err != nil {
		return false
	}
	return match
}

func GeneratePassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}
