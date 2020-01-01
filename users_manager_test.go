package users

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := getTestingUser()
	user.Role = "admin"
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	manager := NewManager(client)
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if client != nil {
			if err := manager.Close(); err != nil {
				t.Fatal(err)
			}
		}
	}()
	storedUser, err := manager.GetUserByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	if storedUser.Email != user.Email {
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
	defer func() {
		err = manager.Close()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	if err := createTwentyUsers(manager); err != nil {
		t.Error(err)
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
