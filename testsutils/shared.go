package testsutils

import (
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/codehell/users"
	"github.com/google/uuid"
)

func GetTestingUser() users.User {
	username, err := users.NewUsername("cazaplanetas")
	if err != nil {
		log.Fatal(err)
	}
	userId, err := users.NewUserID(uuid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	email := "cazaplanetas@gmail.com"
	password := "secret1"
	role := "admin"
	user, err := users.NewUser(userId, username, email, password, role)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func CreateUser() users.User {
	username, _ := users.NewUsername(faker.Username())
	userId, _ := users.NewUserID(uuid.New().String())
	email := faker.Email()
	password := faker.Password()[:12]
	role := "admin"
	user, err := users.NewUser(userId, username, email, password, role)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func CreateTenUsers(client users.UserRepo) error {
	for i := 0; i < 10; i++ {
		user := CreateUser()
		err := client.Store(user)
		if err != nil {
			return err
		}
	}
	return nil
}
