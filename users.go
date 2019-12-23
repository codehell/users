package users

import "time"

type User struct {
	username  string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
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

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) SetCreatedAt(date time.Time) {
	u.createdAt = date
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetUpdatedAt(date time.Time) {
	u.updatedAt = date
}
