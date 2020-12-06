package users

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/alexedwards/argon2id"
	"gopkg.in/yaml.v2"
)

// UserConfig struct that represent users module configuration
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
	UpdatedAt time.Time `json:"update_at"`
}

// NewUser User entity constructor
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

// MarshalJSON interface implementation
func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(userMapper{
		ID:        u.id.Value(),
		Username:  u.username.Value(),
		Email:     u.email,
		Role:      u.role,
		CreatedAt: u.createdAt,
		UpdatedAt: u.updatedAt,
	})
}

// UnmarshalJSON interface implementation
func (u *User) UnmarshalJSON(bytes []byte) error {
	var user userMapper
	if err := json.Unmarshal(bytes, &user); err != nil {
		return err
	}
	username, err := NewUsername(user.Username)
	if err != nil {
		return err
	}
	userID, err := NewUserID(user.ID)
	if err != nil {
		return err
	}
	u.id = userID
	u.username = username
	u.email = user.Email
	u.role = user.Role
	u.createdAt = user.CreatedAt
	u.updatedAt = user.UpdatedAt
	return nil
}

// ID User property getter
func (u User) ID() UserID {
	return u.id
}

// Username User property getter
func (u User) Username() Username {
	return u.username
}

// Email User property getter
func (u User) Email() string {
	return u.email
}

// Password User property getter
func (u User) Password() string {
	return u.password
}

// Role User property getter
func (u User) Role() string {
	return u.role
}

// CreatedAt User property getter
func (u User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt User property getter
func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

// Config users module configuration getter
func Config() (UserConfig, error) {
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

// CheckPassword compare password and hash
func CheckPassword(hash string, password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false
	}
	return match
}

// GeneratePassword generate a argon2id password
func GeneratePassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// UserRepo interface for repositories
type UserRepo interface {
	Store(u User) error
	Search(id string) (User, error)
	SearchByField(value string, field string) (User, error)
	All() ([]User, error)
	Close() error
}
