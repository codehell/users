package users

import "testing"

func TestValidUsername(t *testing.T) {
	client, _ := NewClient()
	manager := NewManager(client)
	user := getTestingUser()
	manager.MinUsernameCharacters = 12
	manager.MaxUsernameCharacters = 12
	if err := manager.CreateUser(&user); err != nil {
		t.Fatal("username should have valid number of chars")
	}
}

func TestInvalidUsername(t *testing.T) {
	client, _ := NewClient()
	manager := NewManager(client)
	user := getTestingUser()
	manager.MinUsernameCharacters = 13
	if err := manager.CreateUser(&user); err != MinUsernameError {
		t.Fatal("username should have less chars than allowed")
	}
	manager.MinUsernameCharacters = 3
	manager.MaxUsernameCharacters = 11
	if err := manager.CreateUser(&user); err != MaxUsernameError {
		t.Fatal("username should have more chars than allowed")
	}
}
