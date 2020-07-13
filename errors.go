package users

import (
	"errors"
	"fmt"
)

var (
	UserPasswordNotMatchError = Error{Code: "userPasswordNotMatch", Err: errors.New("user password not match")}
	UserNotFoundError         = Error{Code: "userNotFound", Err: errors.New("user not found")}
	UserAlreadyExistsError    = Error{Code: "userAlreadyExist", Err: errors.New("user already exist")}
	UserSystemError           = Error{Code: "userSystemError", Err: errors.New("user system error")}
	UserValidationError       = Error{Code: "userValidationError", Err: errors.New("user validation error")}
)

type Error struct {
	Code string
	Err  error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Code, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}
