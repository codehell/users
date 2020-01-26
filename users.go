package users

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var ErrUserAlreadyExists = errors.New("users: user already exists")

type userConfig struct {
	UniqueUsername bool `yaml:"unique_username"`
}

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
	Close() error
	StoreUser(*User) error
	DeleteAll() error
	GetAll() ([]User, error)
	GetUserByEmail(string) (User, error)
	Validator() Validator
	SetValidator(Validator)
}

type Validator func(user *User) error

func StoreUser(u User, c Client) error {
	config, err :=  getConfig()
	if err != nil {
		return err
	}
	validator := c.Validator()
	if err := validator(&u); err != nil {
		return err
	}
	if config.UniqueUsername {
		// intentionally ignored error
		if user, _ := c.GetUserByEmail(u.Email); user.Email != "" {
			return ErrUserAlreadyExists
		}
	}
	if err := c.StoreUser(&u); err != nil {
		return err
	}
	return nil
}

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

func getConfig() (userConfig, error) {
	var config userConfig

	body, err := ioutil.ReadFile("users.config.yaml")
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(body, &config); err != nil {
		return config, err
	}
	return config, nil
}