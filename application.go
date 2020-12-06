package users

import (
	"fmt"
	"log"
)

func StoreUser(userRepo UserRepo, u User) error {
	if err := checkUniqueUsername(userRepo, u); err != nil {
		uniqueUsernameError := fmt.Errorf("unique username error: %w", err)
		log.Println(err)
		return uniqueUsernameError
	}
	if err := userRepo.Store(u); err != nil {
		storeError := fmt.Errorf("error when store a user: %w", err)
		log.Printf("%v", storeError)
		return storeError
	}
	return nil
}

func AllUsers(userRepo UserRepo) ([]User, error) {
	all, err := userRepo.All()
	if err != nil {
		allUsersError := fmt.Errorf("error when get all users: %w", err)
		log.Printf("%v", allUsersError)
		return all, allUsersError
	}
	return all, nil
}

func Search(userRepo UserRepo, id UserID) (User, error) {
	user, err := userRepo.Search(id.Value())
	if err != nil {
		findUserError := fmt.Errorf("error when search user: %w", err)
		log.Printf("%v", findUserError)
		return user, findUserError
	}
	return user, nil
}

func SearchByField(userRepo UserRepo, value string, field string) (User, error) {
	user, err := userRepo.SearchByField(value, field)
	if err != nil {
		findUserError := fmt.Errorf("error when find user by field: %w", err)
		log.Printf("%v", findUserError)
		return user, findUserError
	}
	return user, nil
}

func SignIn(userRepo UserRepo, email string, password string) (User, error) {
	user, err := userRepo.SearchByField(email, "email")
	if err != nil {
		signInUserError := fmt.Errorf("error when signin user: %w", err)
		log.Printf("%v", signInUserError)
		return User{}, signInUserError
	}
	if CheckPassword(user.Password(), password) {
		return user, nil
	}
	return User{}, &UserPasswordNotMatchError
}

// TODO: inject configuration instead just read from file
func checkUniqueUsername(repo UserRepo, user User) error {
	config, err := Config()
	if err != nil {
		configError := fmt.Errorf("getting config error: %w", err)
		log.Println(configError)
		return configError
	}
	if _, err := SearchByField(repo, user.Username().Value(), "username"); config.UniqueUsername && err == nil {
		log.Println(err)
		return &UserAlreadyExistsError
	}
	return nil
}
