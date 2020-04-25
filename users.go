package users

import (
	"encoding/json"
	"errors"
	"github.com/codehell/users/valueobjects"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var ErrUserAlreadyExists = errors.New("users: User already exists")

type UserConfig struct {
	UniqueUsername bool `yaml:"unique_username"`
}

type User struct {
	id        string
	username  valueobjects.Username
	email     string
	password  string
	role      string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id string, username valueobjects.Username, email, password, role string) (User, error) {
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

func (u User) Username() valueobjects.Username {
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
