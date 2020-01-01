package users

import (
	"testing"
)

func TestOkPassword(t *testing.T) {
	user := getTestingUser()
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
		err = manager.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	if !manager.CheckPassword(user, getTestingUser().Password) {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	user := getTestingUser()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	manager := NewManager(client)
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	badPassword := "badPassword"
	if manager.CheckPassword(user, badPassword) {
		t.Error("password should not match")
	}
}
