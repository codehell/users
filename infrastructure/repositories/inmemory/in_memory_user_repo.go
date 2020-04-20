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

func (fuc *UserRepo) StoreUser(u *users.User) error {
	fuc.CreatedUsers = append(fuc.CreatedUsers, *u)
	return nil
}

func (fuc *UserRepo) GetUserByEmail(email string) (users.User, error) {
	for _, v := range fuc.CreatedUsers {
		if v.Email() == email {
			return v, nil
		}
	}
	return users.User{}, errors.New("user not found")
}

func (fuc *UserRepo) GetAll() ([]users.User, error) {
	return fuc.CreatedUsers, nil
}

func (fuc *UserRepo) DeleteAll() error {
	fuc.CreatedUsers = nil
	return nil
}
