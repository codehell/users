package users

import "testing"

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