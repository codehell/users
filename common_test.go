package users

import (
	"github.com/bxcodec/faker/v3"
)

func getTestingUser() User {
	user := User{}
	user.Username = "cazaplanetas"
	user.Email = "cazaplanetas@gmail.com"
	user.Password = "secret1"
	user.Role = "admin"
	return user
}

func createUser() User {
	var user User
	user.Username = faker.Username()
	user.Email = faker.Email()
	user.Password = faker.Password()[:12]
	user.Role = "admin"
	return user
}

func createTwentyUsers(manager *Manager) error {
	for i := 0; i < 20; i++ {
		user := createUser()
		err := manager.CreateUser(&user)
		if err != nil {
			return err
		}
	}
	return nil
}

func createUserClientAndManager(user User) (User, Client, *Manager, error) {
	client, err := NewClient()
	if err != nil {
		return user, nil, nil, err
	}
	manager := NewManager(client)
	err = manager.CreateUser(&user)
	if err != nil {
		return user, nil, nil, err
	}
	return user, client, manager, nil
}
