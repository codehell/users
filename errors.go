package users

import (
	"errors"
	"fmt"
)

var (
	UserPasswordNotMatchError = DomainError{
		Code: "userPasswordNotMatch",
		Err:  errors.New("user password not match"),
	}
	UserNotFoundError = DomainError{
		Code: "userNotFound",
		Err:  errors.New("user not found"),
	}
	UserAlreadyExistsError = DomainError{
		Code: "userAlreadyExist",
		Err:  errors.New("user already exist"),
	}
	UserValidationError = DomainError{
		Code: "userValidationError",
		Err:  errors.New("user validation error"),
	}
)

type DomainError struct {
	Code string
	Err  error
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("code: %s - description: %v", e.Code, e.Err)
}

/*
func (e *DomainError) Unwrap() error {
	return e.Err
}
*/
