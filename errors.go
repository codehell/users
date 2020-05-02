package users

var (
	EmailNotExistError    = Error{Code:"emailNotExist"}
	PassWordNotMatchError = Error{Code: "passwordNotMatch"}
	UserNotExistError = Error{Code: "userNotExist"}
)

type Error struct {
	Code string
}

func (e Error) Error() string {
	return e.Code
}
