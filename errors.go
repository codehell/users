package users

import "fmt"

var (
	UserPassWordNotMatchError = Error{Code: "userPasswordNotMatch"}
	UserNotFoundError         = Error{Code: "userNotFound"}
	UserAlreadyExistsError    = Error{Code: "userAlreadyExist"}
	UserSystemError           = Error{Code: "userSystemError"}
	UserValidationError		  = Error{Code: "userValidationError"}
)

type Error struct {
	Code string
	err error
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.err.Error())
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Wrap(err error) *Error {
	e.err = err
	return e
}
