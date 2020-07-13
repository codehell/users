package users

import "log"

func StoreUser(userRepo UserRepo, u User) error {
	if err := checkUniqueUsername(userRepo, u); err != nil {
		return err
	}
	if err := userRepo.Store(u); err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}

func AllUsers(userRepo UserRepo) ([]User, error) {
	all, err := userRepo.All()
	if err != nil {
		log.Printf("%v", err)
		return all, &UserSystemError
	}
	return all, nil
}

func Find(userRepo UserRepo, id UserID) (User, error) {
	user, err := userRepo.Find(id.Value())
	if err != nil {
		log.Printf("%v", err)
		return user, &UserNotFoundError
	}
	return user, nil
}

func FindByField(userRepo UserRepo, value string, field string) (User, error) {
	user, err := userRepo.FindByField(value, field)
	if err != nil {
		log.Printf("%v", err)
		return user, &UserNotFoundError
	}
	return user, nil
}

func SignIn(userRepo UserRepo, email string, password string) (User, error) {
	user, err := userRepo.FindByField(email, "email")
	if err != nil {
		log.Printf("%v", err)
		return User{}, &UserNotFoundError
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
		log.Println(err)
		return &UserSystemError
	}
	if _, err := FindByField(repo, user.Email(), "email"); config.UniqueUsername && err == nil {
		return &UserAlreadyExistsError
	}
	return nil
}
