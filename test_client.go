package users

import (
	"errors"
)

type FakeUsersClient struct {
	CreatedUsers  []User
	NumberOfUsers int
}

func NewClient() (*FakeUsersClient, error) {
	fakeUsersClient := new(FakeUsersClient)
	fakeUsersClient.NumberOfUsers = 25
	return fakeUsersClient, nil
}

func (*FakeUsersClient) Close() error {
	return nil
}

func (fuc *FakeUsersClient) Create(u *User) error {
	fuc.CreatedUsers = append(fuc.CreatedUsers, *u)
	return nil
}

func (fuc *FakeUsersClient) GetUserByEmail(email string) (User, error) {
	for _, v := range fuc.CreatedUsers {
		if v.email == email {
			return v, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (fuc *FakeUsersClient) GetAll() ([]User, error) {
	return fuc.CreatedUsers, nil
}

func (fuc *FakeUsersClient) DeleteAll() error {
	fuc.CreatedUsers = nil
	return nil
}
