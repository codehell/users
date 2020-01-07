package users

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

func init() {
	body, err := ioutil.ReadFile("users.config.yaml")
	if err != nil {
		log.Println("unable to read file")
	}

	var config config
	if err := yaml.Unmarshal(body, &config); err != nil {
		log.Println(err)
	}
}

var ErrUserAlreadyExists = errors.New("users: user already exists")

type config struct {
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
	Create(*User) error
	Close() error
	GetAll() ([]User, error)
	GetUserByEmail(string) (User, error)
	Validator() Validator
	DeleteAll() error
	SetValidator(Validator)
}

type Validator func(user *User) error

func StoreUser(u User, c Client) error {
	validator := c.Validator()
	if err := validator(&u); err != nil {
		return err
	}
	if user, _ := c.GetUserByEmail(u.Email); user.Email != "" {
		return ErrUserAlreadyExists
	}
	if err := c.Create(&u); err != nil {
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
