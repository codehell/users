package application

import (
	"github.com/codehell/users"
	"github.com/google/martian/log"
)

func StoreUser(userRepo users.UserRepo, u users.User) error {
	// TODO: move finder to listener
	if _, err := FindByField(userRepo, u.Email(), "email"); err == nil {
		return users.UserAlreadyExistsError
	}
	if err := userRepo.Store(u); err != nil {
		log.Errorf("%v", err)
		return err
	}
	return nil
}

func AllUsers(userRepo users.UserRepo) ([]users.User, error) {
	all, err := userRepo.All()
	if err != nil {
		return all, err
	}
	return all, nil
}

func Find(userRepo users.UserRepo, id users.UserID) (users.User, users.DomainError) {
	user, err := userRepo.Find(id.Value())
	if err != nil {
		return user, users.UserNotFoundError
	}
	return user, nil
}

func FindByField(userRepo users.UserRepo, value string, field string) (users.User, users.DomainError) {
	user, err := userRepo.FindByField(value, field)
	if err != nil {
		return user, users.UserNotFoundError
	}
	return user, nil
}

func SignIn(userRepo users.UserRepo, email string, password string) (users.User, users.DomainError) {
	user, err := userRepo.FindByField(email, "email")
	if err != nil {
		return users.User{}, users.UserNotFoundError
	}
	if users.CheckPassword(user.Password(), password) {
		return user, nil
	}
	return users.User{}, users.UserPasswordNotMatchError
}
