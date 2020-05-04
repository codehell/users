package users

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type UserConfig struct {
	UniqueUsername bool `yaml:"unique_username"`
}

// User entity
type User struct {
	id        UserID
	username  Username
	email     string
	password  string
	role      string
	createdAt time.Time
	updatedAt time.Time
}

type userMapper struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewUser(id UserID, username Username, email, password, role string) (User, error) {
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
	user := userMapper{
		u.id.Value(),
		u.username.Value(),
		u.email,
		u.role,
		u.createdAt,
		u.updatedAt,
	}
	return json.Marshal(user)
}

func (u *User) UnmarshalJSON(bytes []byte) error {
	user := userMapper{}
	if err := json.Unmarshal(bytes, &user); err != nil {
		return err
	}
	username, err := NewUsername(user.Username)
	if err != nil {
		return err
	}
	userId, err := NewUserId(user.ID)
	u.id = userId
	u.username = username
	u.email = user.Email
	u.role = user.Role
	u.createdAt = user.CreatedAt
	u.updatedAt = user.UpdateAt
	return nil
}

func (u User) Id() UserID {
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
	Store(u User) error
	Find(id string) (User, error)
	FindByField(value string, field string) (User, error)
	GetAll() ([]User, error)
}
