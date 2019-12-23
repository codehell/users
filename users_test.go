package users

import (
	"github.com/bxcodec/faker/v3"
	"testing"
)

func getTestingUser() User {
	user := User{}
	user.SetUsername("cazaplanetas")
	user.SetEmail("cazaplanetas@gmail.com")
	user.SetPassword("secret")
	return user
}

func createUser() User{
	var user User
	user.SetUsername(faker.Username())
	user.SetEmail(faker.Email())
	user.SetPassword(faker.Password())
	return user
}

func createTwentyUsers(manager *UserManager) error {
	for i := 0; i < 20; i++ {
		user := createUser()
		err := manager.CreateUser(&user)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestCreateUser(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func () {
		err = client.CloseClient()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	manager := NewManager(client)
	user := getTestingUser()
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	storedUser, err := manager.GetUserByEmail(user.Email())
	if err != nil {
		t.Fatal(err)
	}
	if storedUser.Username() != user.Username() {
		t.Error("The user was not the expected user")
	}
	if err = client.deleteAll(); err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func() {
		err = client.CloseClient()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	manager := NewManager(client)
	if err := createTwentyUsers(manager); err != nil {
		t.Fatal(err)
	}
	myUsers, err := manager.GetUsers()
	if err != nil {
		t.Fatal(err)
	}

	if len(myUsers) != 20 {
		t.Errorf("got %d users, expect %d users", len(myUsers), 20)
	}

	if err = client.deleteAll(); err != nil {
		t.Fatal(err)
	}
}
