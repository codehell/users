package users

import (
	"testing"
)

// TODO: Estos test estan provando el validador por defecto
// Realmente deberia hacer un test unitario del validador
// Y otros test mas genericos para el manager

func TestValidUsername(t *testing.T) {
	client, _ := NewClient()
	manager := NewManager(client)
	user := getTestingUser()
	if err := manager.CreateUser(&user); err != nil {
		t.Error("username should have valid number of chars")
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
		t.Error("username should have more chars than allowed")
	}
}

func TestRoleCanNotBeEmpty(t *testing.T) {
	client, _ := NewClient()
	manager := NewManager(client)
	user := getTestingUser()
	user.SetRole("")
	if err := manager.CreateUser(&user); err == nil {
		t.Error("user manager create a user with empty role")
	}
}
