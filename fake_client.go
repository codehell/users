package users

import (
	"errors"
	"github.com/bxcodec/faker/v3"
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

func (*FakeUsersClient) CloseClient() error {
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
	for i := 0; i <= fuc.NumberOfUsers; i++ {
		user := createUser()
		fuc.CreatedUsers = append(fuc.CreatedUsers, user)
	}
	return fuc.CreatedUsers, nil
}

func (fuc *FakeUsersClient) deleteAll() error {
	fuc.CreatedUsers = nil
	return nil
}

func createUser() User{
	password, err := generatePassword(faker.Password())
	if err != nil {
		password = ""
	}
	return User{
		username: faker.Username(),
		email:    faker.Email(),
		password: password,
	}
}
