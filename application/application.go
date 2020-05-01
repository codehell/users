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
	return userRepo.Find(id)
}

func FindField(userRepo users.UserRepo, value string, field string) (users.User, error) {
	return userRepo.FindField(value, field)
}
