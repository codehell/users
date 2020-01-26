package users_test

import (
	"github.com/codehell/users"
	"testing"
)

func TestDefaultValidatorError(t *testing.T) {
	user := getTestingUser()
	user.Role = ""
	if err := users.DefaultValidator(&user); err != users.ErrInvalidRole {
		t.Errorf("expect %v, got %v", users.ErrInvalidRole, err)
	}
	user.Password = "6?ZM]jnM:T3AgSMu#$O$w\"31Z~Rx?lBMA"
	if err := users.DefaultValidator(&user); err != users.ErrMaxPassword {
		t.Errorf("expect %v, got %v", users.ErrMaxPassword, err)
	}
	user.Password = "6?ZM]"
	if err := users.DefaultValidator(&user); err != users.ErrMinPassword {
		t.Errorf("expect %v, got %v", users.ErrMinPassword, err)
	}
	user.Username = "supercalifragilisticoespialidoso"
	if err := users.DefaultValidator(&user); err != users.ErrMaxUsername {
		t.Errorf("expect %v, got %v", users.ErrMaxUsername, err)
	}
	user.Username = "caz"
	if err := users.DefaultValidator(&user); err != users.ErrMinUsername {
		t.Errorf("expect %v, got %v", users.ErrMinUsername, err)
	}
}
