package firestore

import (
	"github.com/codehell/users"
	"testing"
)

func getTestingUser() users.User {
	user := users.User{}
	user.SetUsername("cazaplanetas")
	user.SetEmail("cazaplanetas@gmail.com")
	user.SetPassword("secret")
	return user
}

func TestCreateUser(t *testing.T) {
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = client.CloseClient()
		if err != nil {
			t.Fatal(err)
		}
	}()
	manager := users.NewManager(client)
	user := getTestingUser()
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	storedUser, err := manager.GetUserByEmail(user.Email())
	if err != nil {
		t.Fatal(err)
	}
	if storedUser != user {
		t.Errorf("The user %v was not the expected user %v", storedUser, user)
	}
	if err = client.deleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func() {
		err = client.CloseClient()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	manager := users.NewManager(client)
	user := getTestingUser()
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	myUsers, err := manager.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
	if myUsers[0] != user {
		t.Errorf("got %v, expect %v", myUsers[0], user)
	}
	if err = client.deleteAll(); err != nil {
		t.Fatal(err)
	}
}
