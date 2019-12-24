package users

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"time"
)

var MinUsernameError = errors.New("username is too sort")
var MaxUsernameError = errors.New("username is too long")
var MinPasswordError = errors.New("too sort password")
var MaxPasswordError = errors.New("too long password")


type UserManager struct {
	client Client
	MinUsernameCharacters int
	MaxUsernameCharacters int
	MinPasswordCharacters int
	MaxPasswordCharacters int
}

type Client interface {
	Create(*User) error
	Close() error
	GetUserByEmail(string) (User, error)
	GetAll() ([]User, error)
	DeleteAll() error
}

func NewManager(c Client) *UserManager {
	um := new(UserManager)
	um.client = c
	um.MinUsernameCharacters = 3
	um.MaxUsernameCharacters = 16
	um.MinPasswordCharacters = 6
	um.MaxPasswordCharacters = 32
	return um
}

func (um *UserManager) CreateUser(u *User) error {
	if err := validateUser(um, u); err != nil {
		return err
	}
	password, err := generatePassword(u.password)
	if err != nil {
		return err
	}
	u.password = password
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return um.client.Create(u)
}

func (um *UserManager) GetUserByEmail(email string) (User, error) {
	return um.client.GetUserByEmail(email)
}

func (um *UserManager) GetUsers() ([]User, error) {
	return um.client.GetAll()
}

func (um *UserManager) Close() error {
	if um.client != nil {
		return um.client.Close()
	}
	return errors.New("client is nil")
}

func generatePassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func validateUser(um *UserManager, u *User) error {
	usernameLen := len(u.Username())
	if usernameLen < um.MinUsernameCharacters {
		return MinUsernameError
	}
	if usernameLen > um.MaxUsernameCharacters {
		return MaxUsernameError
	}
	if usernameLen < um.MinPasswordCharacters {
		return MinPasswordError
	}
	if usernameLen > um.MaxPasswordCharacters {
		return MaxPasswordError
	}
	return nil
}