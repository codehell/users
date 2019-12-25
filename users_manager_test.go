package users

import "testing"

func TestValidUsername(t *testing.T) {
	client, _ := NewClient()
	manager := NewManager(client)
	user := getTestingUser()
	if err := manager.CreateUser(&user); err != nil {
		t.Fatal("username should have valid number of chars")
	}
}

func TestInvalidUsername(t *testing.T) {
	client, _ := NewClient()
	manager := NewManager(client)
	user := getTestingUser()
	user.SetUsername("caz")
	if err := manager.CreateUser(&user); err != MinUsernameError {
		t.Fatal("username should have less chars than allowed")
	}
	user.SetUsername("supercalifragilistospialidoso")
	if err := manager.CreateUser(&user); err != MaxUsernameError {
		t.Fatal("username should have more chars than allowed")
	}
}
