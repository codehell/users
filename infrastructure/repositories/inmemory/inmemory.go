package inmemory

import (
	"errors"
	"github.com/codehell/users"
)

type UserRepo struct {
	CreatedUsers []users.User
}

func NewRepo() (*UserRepo, error) {
	fakeUsersClient := new(UserRepo)
	return fakeUsersClient, nil
}

func (*UserRepo) Close() error {
	return nil
}

func (fuc *UserRepo) Store(u users.User) error {
	_, err := fuc.FindByField(u.Email(), "email")
	if err == nil {
		return users.UserAlreadyExistsError
	}
	_, err = fuc.FindByField(u.Username().Value(), "username")
	if err == nil {
		return users.UserAlreadyExistsError
	}
	fuc.CreatedUsers = append(fuc.CreatedUsers, u)
	return nil
}

func (fuc *UserRepo) Find(id string) (users.User, error) {
	for _, u := range fuc.CreatedUsers {
		if u.Id().Value() == id {
			return u, nil
		}
	}
	return users.User{}, users.UserNotFoundError
}

func (fuc *UserRepo) FindByField(value string, field string) (users.User, error) {
	if field != "username" && field != "email" {
		return users.User{}, errors.New("bad field")
	}
	for _, u := range fuc.CreatedUsers {
		userMap := map[string]string{"username": u.Username().Value(), "email": u.Email()}
		if userMap[field] == value {
			return u, nil
		}
	}
	return users.User{}, users.UserNotFoundError
}

func (fuc *UserRepo) GetAll() ([]users.User, error) {
	return fuc.CreatedUsers, nil
}

func (fuc *UserRepo) DeleteAll() error {
	fuc.CreatedUsers = nil
	return nil
}
