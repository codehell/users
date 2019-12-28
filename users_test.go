package users

import (
	"testing"
)

func TestOkPassword(t *testing.T) {
	testingUser := getTestingUser()
	user, _, manager, err := createUserClientAndManager(testingUser)
	defer func() {
		err = manager.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	if !user.CheckPassword(testingUser.Password) {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	testingUser := getTestingUser()
	user, _, manager, err := createUserClientAndManager(testingUser)
	defer func() {
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
