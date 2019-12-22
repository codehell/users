package users

import "testing"

func getTestingUser() User {
	return User{
		username: "cazaplanetas",
		email:    "cazaplanetas@gmail.com",
		password: "secret",
	}
}

func TestOkPassword(t *testing.T) {
	user := getTestingUser()
	originalPassword := user.password
	hash, err := generatePassword(user.password)
	user.password = hash
	if err != nil {
		t.Error(err)
	}
	if !user.CheckPassword(originalPassword) {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	user := getTestingUser()
	badPassword := "badPassword"
	hash, err := generatePassword(user.password)
	user.password = hash
	if err != nil {
		t.Error(err)
	}
	if user.CheckPassword(badPassword) {
		t.Error("password should not match")
	}
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
	user := User{"codehell", "cazaplanetas@gmail.com", "secret"}
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	storedUser, err := manager.GetUserByEmail(user.email)
	if err != nil {
		t.Fatal(err)
	}
	if storedUser != user {
		t.Error("The user was not the expected user")
	}
	if err = client.deleteAll(); err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error("can not create client")
	}
	defer func () {
		err = client.CloseClient()
		if err != nil {
			t.Error("can not close client")
		}
	}()
	manager := NewManager(client)
	_, err = manager.GetUsers()
	if err != nil {
		t.Error(err)
	}
	if err = client.deleteAll(); err != nil {
		t.Error(err)
	}
}
