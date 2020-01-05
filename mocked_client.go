package users

import (
	"errors"
)

type FakeUsersClient struct {
	CreatedUsers  []User
	validator Validator
}

func NewClient() (*FakeUsersClient, error) {
	fakeUsersClient := new(FakeUsersClient)
	fakeUsersClient.validator = DefaultValidator
	return fakeUsersClient, nil
}

func (*FakeUsersClient) Close() error {
	return nil
}

func (fuc *FakeUsersClient) Create(u *User) error {
	if err := fuc.validator(u); err != nil {
		return err
	}
	u.ID = u.Email
	var err error
	if u.Password, err = GeneratePassword(u.Password); err != nil {
		return err
	}
	fuc.CreatedUsers = append(fuc.CreatedUsers, *u)
	return nil
}

func (fuc *FakeUsersClient) GetUserByEmail(email string) (User, error) {
	for _, v := range fuc.CreatedUsers {
		if v.Email == email {
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

func (fuc *FakeUsersClient) SetValidator(v Validator) {
	fuc.validator = v
}