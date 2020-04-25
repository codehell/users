package application

import (
	"github.com/codehell/users"
)

func StoreUser(u users.User, userRepo users.UserRepo) error {
	config, err :=  users.GetConfig()
	if err != nil {
		return err
	}
	if config.UniqueUsername {
		// intentionally ignored error
		if user, _ := userRepo.GetUserByEmail(u.Email()); user.Email() != "" {
			return users.ErrUserAlreadyExists
		}
	}
	if err := userRepo.StoreUser(&u); err != nil {
		return err
	}
	return nil
}

func GetAll(userRepo users.UserRepo) ([]users.User, error) {
	all, err := userRepo.GetAll()
	if err != nil {
		return all, err
	}
	return all, nil
}
