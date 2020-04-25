package shared

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codehell/users"
	"github.com/codehell/users/valueobjects"
	"github.com/google/uuid"
	"log"
)

func GetTestingUser() users.User {
	username, err := valueobjects.NewUsername("cazaplanetas")
	if err != nil {
		log.Fatal(err)
	}
	email := "cazaplanetas@gmail.com"
	password := "secret1"
	role := "admin"
	user, err := users.NewUser(uuid.New().String(), username, email, password, role)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func CreateUser() users.User {
	username, _ := valueobjects.NewUsername(faker.Username())
	email := faker.Email()
	password := faker.Password()[:12]
	role := "admin"
	user, err := users.NewUser(uuid.New().String(), username, email, password, role)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func CreateTwentyUsers(client users.UserRepo) error {
	for i := 0; i < 20; i++ {
		user := CreateUser()
		err := client.StoreUser(&user)
		if err != nil {
			return err
		}
	}
	return nil
}
