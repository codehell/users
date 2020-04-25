package users

import (
	"encoding/json"
	"errors"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var ErrUserAlreadyExists = errors.New("users: User already exists")

type UserConfig struct {
	UniqueUsername bool `yaml:"unique_username"`
}

// User entity
type User struct {
	id        string
	username  Username
	email     string
	password  string
	role      string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id string, username Username, email, password, role string) (User, error) {
	cryptPassword, err := GeneratePassword(password)
	if err != nil {
		return User{}, err
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	return User{
		id:        id,
		username:  username,
		email:     email,
		password:  cryptPassword,
		role:      role,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (u User) MarshalJSON() ([]byte, error) {
	user := struct {
		ID        string    `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdateAt  time.Time `json:"update_at"`
	}{
		u.id,
		u.username.Value(),
		u.email,
		u.role,
		u.createdAt,
		u.updatedAt,
	}
	return json.Marshal(user)
}

func (u User) Id() string {
	return u.id
}

func (u User) Username() Username {
	return u.username
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) Role() string {
	return u.role
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

func GetConfig() (UserConfig, error) {
	var config UserConfig
	body, err := ioutil.ReadFile("users.config.yaml")
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(body, &config); err != nil {
		return config, err
	}
	return config, nil
}

// UserRepo interface for repositories
type UserRepo interface {
	Close() error
	StoreUser(*User) error
	DeleteAll() error
	GetAll() ([]User, error)
	GetUserByEmail(string) (User, error)
}

// UserID value object
type UserID struct {
	value string
}

// Username value object
type Username struct {
	value string
}

func NewUsername(name string) (Username, error) {
	validate := validator.New()
	// Los errores de la libreria de validación pueden usarse
	// desde el momento que añado la libreria al dominio
	err := validate.Var(name, "min=5,max=64")
	if err != nil {
		return Username{}, err
	}
	return Username{name}, nil
}

func (un Username) validate() error {
	return nil
}

func (un Username) isEqualTo(username Username) bool {
	return un.value == username.Value()
}

func (un Username) Value() string {
	return un.value
}
