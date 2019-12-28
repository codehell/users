package users

import (
	"github.com/bxcodec/faker/v3"
	"testing"
)

func getTestingUser() User {
	user := User{}
	user.SetUsername("cazaplanetas")
	user.SetEmail("cazaplanetas@gmail.com")
	user.SetPassword("secret1")
	user.SetRole("admin")
	return user
}

func createUser() User{
	var user User
	user.SetUsername(faker.Username())
	user.SetEmail(faker.Email())
	user.SetPassword(faker.Password())
	user.SetRole("admin")
	return user
}

func createTwentyUsers(manager *Manager) error {
	for i := 0; i < 20; i++ {
		user := createUser()
		err := manager.CreateUser(&user)
		if err != nil {
			return err
		}
	}
	return nil
}

func createUserClientAndManager(user User) (User, Client, *Manager, error) {
	client, err := NewClient()
	if err != nil {
		return user, nil, nil, err
	}
	manager := NewManager(client)
	err = manager.CreateUser(&user)
	if err != nil {
		return user, nil, nil, err
	}
	return user, client, manager, nil
}

func TestCreateUser(t *testing.T) {
	testingUser := getTestingUser()
	testingUser.SetRole("admin")
	user, client, manager, err := createUserClientAndManager(testingUser)
	defer func () {
		err = manager.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
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
	if err = client.DeleteAll(); err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal("can not create client")
	}
	manager := NewManager(client)
	defer func () {
		err = manager.Close()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
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

	if err = client.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestOkPassword(t *testing.T) {
	testingUser := getTestingUser()
	user, _, manager, err := createUserClientAndManager(testingUser)
	defer func () {
		err = manager.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	if !user.CheckPassword(testingUser.Password()) {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	testingUser := getTestingUser()
	user, _, manager, err := createUserClientAndManager(testingUser)
	defer func () {
		err = manager.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	badPassword := "badPassword"
	if err != nil {
		t.Error(err)
	}
	if user.CheckPassword(badPassword) {
		t.Error("password should not match")
	}
}