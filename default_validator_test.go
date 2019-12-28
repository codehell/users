package users

import "testing"

func TestDefaultValidatorError(t *testing.T) {
	user := getTestingUser()
	user.SetRole("")
	if err := defaultValidator(user); err != InvalidRoleError {
		t.Errorf("expect %v, got %v", InvalidRoleError, err)
	}
	user.SetPassword("6?ZM]jnM:T3AgSMu#$O$w\"31Z~Rx?lBMA")
	if err := defaultValidator(user); err != MaxPasswordError {
		t.Errorf("expect %v, got %v", MaxPasswordError, err)
	}
	user.SetPassword("6?ZM]")
	if err := defaultValidator(user); err != MinPasswordError {
		t.Errorf("expect %v, got %v", MinPasswordError, err)
	}
	user.SetUsername("supercalifragilisticoespialidoso")
	if err := defaultValidator(user); err != MaxUsernameError {
		t.Errorf("expect %v, got %v", MaxUsernameError, err)
	}
	user.SetUsername("caz")
	if err := defaultValidator(user); err != MinUsernameError {
		t.Errorf("expect %v, got %v", MinUsernameError, err)
	}
}