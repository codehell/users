package application

import (
	"github.com/codehell/users"
)

func StoreUser(u users.User, userRepo users.UserRepo) error {
	if err := userRepo.Store(u); err != nil {
		return err
	}
	return nil
}

func AllUsers(userRepo users.UserRepo) ([]users.User, error) {
	all, err := userRepo.GetAll()
	if err != nil {
		return all, err
	}
	return all, nil
}

func Find(userRepo users.UserRepo, id string) (users.User, error) {
	user, err := userRepo.Find(id)
	if err != nil {
		return user, users.UserNotExistError
	}
	return user, nil
}

func FindByField(userRepo users.UserRepo, value string, field string) (users.User, error) {
	user, err := userRepo.FindByField(value, field)
	if err != nil {
		return user, users.UserNotExistError
	}
	return user, nil
}

func SignIn(userRepo users.UserRepo, email string, password string) (users.User, error) {
	user, err := userRepo.FindByField(email, "email")
	if err != nil {
		return users.User{}, users.EmailNotExistError
	}
	if users.CheckPassword(user.Password(), password) {
		return user, nil
	}
	return users.User{}, users.PassWordNotMatchError
}
