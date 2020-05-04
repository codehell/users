package application

import (
	"github.com/codehell/users"
)

func StoreUser(userRepo users.UserRepo, u users.User) error {
	return userRepo.Store(u)
}

func AllUsers(userRepo users.UserRepo) ([]users.User, error) {
	all, err := userRepo.GetAll()
	if err != nil {
		return all, err
	}
	return all, nil
}

func Find(userRepo users.UserRepo, id users.UserID) (users.User, error) {
	user, err := userRepo.Find(id.Value())
	if err != nil {
		return user, err
	}
	return user, nil
}

func FindByField(userRepo users.UserRepo, value string, field string) (users.User, error) {
	user, err := userRepo.FindByField(value, field)
	if err != nil {
		return user, users.UserNotFoundError
	}
	return user, nil
}

func SignIn(userRepo users.UserRepo, email string, password string) (users.User, error) {
	user, err := userRepo.FindByField(email, "email")
	if err != nil {
		return users.User{}, users.UserNotFoundError
	}
	if users.CheckPassword(user.Password(), password) {
		return user, nil
	}
	return users.User{}, users.UserPassWordNotMatchError
}
