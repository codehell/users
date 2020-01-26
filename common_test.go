package users_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codehell/users"
)

func getTestingUser() users.User {
	user := users.User{}
	user.Username = "cazaplanetas"
	user.Email = "cazaplanetas@gmail.com"
	user.Password = "secret1"
	user.Role = "admin"
	return user
}

func createUser() users.User {
	var user users.User
	user.Username = faker.Username()
	user.Email = faker.Email()
	user.Password = faker.Password()[:12]
	user.Role = "admin"
	return user
}

func createTwentyUsers(client users.Client) error {
	for i := 0; i < 20; i++ {
		user := createUser()
		err := client.StoreUser(&user)
		if err != nil {
			return err
		}
	}
	return nil
}
