package users

import (
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
