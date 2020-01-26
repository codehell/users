package mocked_client

import (
	"errors"
	"github.com/codehell/users"
)

type FakeUsersClient struct {
	CreatedUsers []users.User
	validator    users.Validator
}

func NewClient() (*FakeUsersClient, error) {
	fakeUsersClient := new(FakeUsersClient)
	fakeUsersClient.validator = users.DefaultValidator
	return fakeUsersClient, nil
}

func (*FakeUsersClient) Close() error {
	return nil
}

func (fuc *FakeUsersClient) StoreUser(u *users.User) error {
	if err := fuc.validator(u); err != nil {
		return err
	}
	u.ID = u.Email
	var err error
	if u.Password, err = users.GeneratePassword(u.Password); err != nil {
		return err
	}
	fuc.CreatedUsers = append(fuc.CreatedUsers, *u)
	return nil
}

func (fuc *FakeUsersClient) GetUserByEmail(email string) (users.User, error) {
	for _, v := range fuc.CreatedUsers {
		if v.Email == email {
			return v, nil
		}
	}
	return users.User{}, errors.New("user not found")
}

func (fuc *FakeUsersClient) GetAll() ([]users.User, error) {
	return fuc.CreatedUsers, nil
}

func (fuc *FakeUsersClient) DeleteAll() error {
	fuc.CreatedUsers = nil
	return nil
}

func (fuc *FakeUsersClient) SetValidator(v users.Validator) {
	fuc.validator = v
}

func (fuc *FakeUsersClient) Validator() users.Validator {
	return fuc.validator
}