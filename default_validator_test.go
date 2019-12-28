package users

import "testing"

func TestDefaultValidatorError(t *testing.T) {
	user := getTestingUser()
	user.Role = ""
	if err := defaultValidator(user); err != InvalidRoleError {
		t.Errorf("expect %v, got %v", InvalidRoleError, err)
	}
	user.Password = "6?ZM]jnM:T3AgSMu#$O$w\"31Z~Rx?lBMA"
	if err := defaultValidator(user); err != MaxPasswordError {
		t.Errorf("expect %v, got %v", MaxPasswordError, err)
	}
	user.Password = "6?ZM]"
	if err := defaultValidator(user); err != MinPasswordError {
		t.Errorf("expect %v, got %v", MinPasswordError, err)
	}
	user.Username = "supercalifragilisticoespialidoso"
	if err := defaultValidator(user); err != MaxUsernameError {
		t.Errorf("expect %v, got %v", MaxUsernameError, err)
	}
	user.Username = "caz"
	if err := defaultValidator(user); err != MinUsernameError {
		t.Errorf("expect %v, got %v", MinUsernameError, err)
	}
}
