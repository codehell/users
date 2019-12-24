package users

import (
	"github.com/alexedwards/argon2id"
	"time"
)

type User struct {
	username  string
	email     string
	password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) CheckPassword(password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, u.password)
	if err != nil {
		return false
	}
	return match
}