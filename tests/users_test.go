package tests

import (
	"encoding/json"
	"github.com/codehell/users"
	"strings"
	"testing"
)

func TestOkPassword(t *testing.T) {
	user := GetTestingUser()

	if !users.CheckPassword(user.Password(), "secret1") {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	user := GetTestingUser()

	badPassword := "badPassword"
	if users.CheckPassword(user.Password(), badPassword) {
		t.Error("password should not match")
	}
}

func TestMarshalUser(t *testing.T)  {
	user := GetTestingUser()
	userJson, err := json.Marshal(user)
	if err != nil {
		t.Error("can not marshal user struct")
	}
	if !strings.Contains(string(userJson), "cazaplanetas@gmail.com") {
		t.Errorf("json: %s is incorrect", userJson)
	}
}

func TestUnmarshalUser(t *testing.T) {
	user := GetTestingUser()
	userJson, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	userStruct := users.User{}
	err = json.Unmarshal(userJson, &userStruct); if err != nil {
		t.Fatal(err)
	}
	if userStruct.Email() != "cazaplanetas@gmail.com" {
		t.Errorf("expect user struct with email cazaplanetas@gmail.com, got: %v", userStruct)
	}
}
