package users

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	testingUser := getTestingUser()
	testingUser.Role = "admin"
	user, client, manager, err := createUserClientAndManager(testingUser)
	defer func() {
		err = manager.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}
	storedUser, err := manager.GetUserByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	if storedUser.Username != user.Username {
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
