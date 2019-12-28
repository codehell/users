package users

import (
	"github.com/alexedwards/argon2id"
	"time"
)

type User struct {
	ID        interface{}
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CheckPassword(password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, u.Password)
	if err != nil {
		return false
	}
	return match
}
