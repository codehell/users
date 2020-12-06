package inmemoryrepo

import (
	"errors"
	"fmt"
	"github.com/codehell/users"
)

type UserRepo struct {
	CreatedUsers []users.User
}

func New() (*UserRepo, error) {
	fakeUsersClient := new(UserRepo)
	return fakeUsersClient, nil
}

func (*UserRepo) Close() error {
	return nil
}

func (r *UserRepo) Store(u users.User) error {
	r.CreatedUsers = append(r.CreatedUsers, u)
	return nil
}

func (r *UserRepo) Search(id string) (users.User, error) {
	for _, u := range r.CreatedUsers {
		if u.ID().Value() == id {
			return u, nil
		}
	}
	searchError := fmt.Errorf("user with id: %s does not exist: %w", id, &users.UserNotFoundError)
	return users.User{}, searchError
}

func (r *UserRepo) SearchByField(value string, field string) (users.User, error) {
	if field != "username" && field != "email" {
		return users.User{}, errors.New("wrong field to find")
	}
	for _, u := range r.CreatedUsers {
		userMap := map[string]string{"username": u.Username().Value(), "email": u.Email()}
		if userMap[field] == value {
			return u, nil
		}
	}
	return users.User{}, fmt.Errorf("can not found user by field %s and %s value", field, value)
}

func (r *UserRepo) All() ([]users.User, error) {
	return r.CreatedUsers, nil
}

func (r *UserRepo) DeleteAll() error {
	r.CreatedUsers = nil
	return nil
}
