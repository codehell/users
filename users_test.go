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
	err = client.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if client != nil {
			if err := client.Close(); err != nil {
				t.Fatal(err)
			}
		}
	}()
	storedUser, err := client.GetUserByEmail(user.Email)
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

func TestStoreUser(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	user := getTestingUser()
	if err := StoreUser(user, client); err != nil {
		t.Error(err)
	}
	if err := StoreUser(user, client); err != ErrUserAlreadyExists {
		t.Error(err)
	}
	if err = client.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	if err := createTwentyUsers(client); err != nil {
		t.Error(err)
	}
	myUsers, err := client.GetAll()
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
	user := getTestingUser()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	err = client.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	if !CheckPassword(user, getTestingUser().Password) {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	user := getTestingUser()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	err = client.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	badPassword := "badPassword"
	if CheckPassword(user, badPassword) {
		t.Error("password should not match")
	}
}
