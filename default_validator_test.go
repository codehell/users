package users

import "testing"

func TestDefaultValidatorError(t *testing.T) {
	user := getTestingUser()
	user.Role = ""
	if err := DefaultValidator(&user); err != ErrInvalidRole {
		t.Errorf("expect %v, got %v", ErrInvalidRole, err)
	}
	user.Password = "6?ZM]jnM:T3AgSMu#$O$w\"31Z~Rx?lBMA"
	if err := DefaultValidator(&user); err != ErrMaxPassword {
		t.Errorf("expect %v, got %v", ErrMaxPassword, err)
	}
	user.Password = "6?ZM]"
	if err := DefaultValidator(&user); err != ErrMinPassword {
		t.Errorf("expect %v, got %v", ErrMinPassword, err)
	}
	user.Username = "supercalifragilisticoespialidoso"
	if err := DefaultValidator(&user); err != ErrMaxUsername {
		t.Errorf("expect %v, got %v", ErrMaxUsername, err)
	}
	user.Username = "caz"
	if err := DefaultValidator(&user); err != ErrMinUsername {
		t.Errorf("expect %v, got %v", ErrMinUsername, err)
	}
}
