package users

import (
	"errors"
	"fmt"
	"github.com/alexedwards/argon2id"
)

var MinUsernameCharacters = 3
var MinPasswordCharacters = 6

type UserManager struct {
	client UsersClient
}

type UsersClient interface {
	Create(*User) error
	GetUserByEmail(string) (User, error)
	GetAll() ([]User, error)
}

func NewManager(uc UsersClient) *UserManager {
	um := new(UserManager)
	um.client = uc
	return um
}

func (um *UserManager) CreateUser(u *User) error {
	if len(u.username) <= MinUsernameCharacters {
		return errors.New(fmt.Sprintf("username must have more than %d characters", MinUsernameCharacters))
	}
	if len(u.username) <= MinPasswordCharacters {
		return errors.New(fmt.Sprintf("username must have more than %d characters", MinPasswordCharacters))
	}
	password, err := generatePassword(u.password)
	u.password = password
	if err != nil {
		return err
	}
	return um.client.Create(u)
}

func (um *UserManager) GetUserByEmail(email string) (User, error) {
	return um.client.GetUserByEmail(email)
}

func (um *UserManager) GetUsers() ([]User, error) {
	return um.client.GetAll()
}

func (u *User) CheckPassword(password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, u.password)
	if err != nil {
		return false
	}
	return match
}

func generatePassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}
