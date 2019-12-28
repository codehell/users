package users

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"time"
)

type Manager struct {
	client    Client
	validator Validator
}

type Client interface {
	Create(*User) error
	Close() error
	GetUserByEmail(string) (User, error)
	GetAll() ([]User, error)
	DeleteAll() error
}

type Validator func(user User) error

func NewManager(c Client) *Manager {
	um := new(Manager)
	um.client = c
	um.validator = defaultValidator
	return um
}

func (um *Manager) CreateUser(u *User) error {
	if err := defaultValidator(*u); err != nil {
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

func (um *Manager) GetUserByEmail(email string) (User, error) {
	return um.client.GetUserByEmail(email)
}

func (um *Manager) GetUsers() ([]User, error) {
	return um.client.GetAll()
}

func (um *Manager) Close() error {
	if um.client != nil {
		return um.client.Close()
	}
	return errors.New("client is nil")
}

func (um *Manager) SetValidator(val Validator) {
	um.validator = val
}

func generatePassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}
