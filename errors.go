package users

import (
	"errors"
	"fmt"
)

var (
	UserPasswordNotMatchError = Error{code: "userPasswordNotMatch", err: errors.New("user password not match")}
	UserNotFoundError         = Error{code: "userNotFound", err: errors.New("user not found")}
	UserAlreadyExistsError    = Error{code: "userAlreadyExist", err: errors.New("user already exist")}
	UserSystemError           = Error{code: "userSystemError", err: errors.New("user system error")}
	UserValidationError       = Error{code: "userValidationError", err: errors.New("user validation error")}
)

type DomainError interface {
	Error() string
	Code() string
	Unwrap() error
}

type Error struct {
	code string
	err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.err.Error())
}

func (e Error) Unwrap() error {
	return e.err
}

func (e Error) Code() string {
	return e.code
}
