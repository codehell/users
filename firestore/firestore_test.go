package firestore

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codehell/users"
	"testing"
)

func getTestingUser() users.User {
	user := users.User{}
	user.Username = "cazaplanetas"
	user.Email = "cazaplanetas@gmail.com"
	user.Password = "secret"
	user.Role = "user"
	return user
}

func createUser() users.User {
	var user users.User
	user.Username = faker.Username()
	user.Email = faker.Email()
	user.Password = faker.Password()[:12]
	user.Role = "user"
	return user
}

func createTwentyUsers(client Client) error {
	for i := 0; i < 20; i++ {
		user := createUser()
		err := client.Create(&user)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestCreateUser(t *testing.T) {
	user := getTestingUser()
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Create(&user)
	if err != nil {
		t.Fatal()
	}
	defer func() {
		if err := client.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	storedUser, err := client.GetUserByEmail(user.Email)
	if err != nil {
		t.Fatal(err)
	}
	if storedUser.Username != user.Username {
		t.Error("The user was not the expected user")
	}

	if err = client.DeleteAll(); err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Log("can not close client")
		}
	}()
	if err := createTwentyUsers(*client); err != nil {
		t.Fatal(err)
	}
	myUsers, err := client.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(myUsers) != 20 {
		t.Errorf("got %d users, expect %d users", len(myUsers), 20)
	}

	if err = client.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestOkPassword(t *testing.T) {
	user := getTestingUser()
	password := user.Password
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Log("can not close client")
		}
	}()
	err = client.Create(&user)
	if err != nil {
		t.Fatal()
	}
	if !users.CheckPassword(user, password) {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	user := getTestingUser()
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Log("can not close client")
		}
	}()
	err = client.Create(&user)
	if err != nil {
		t.Fatal()
	}
	badPassword := "badPassword"
	if err != nil {
		t.Error(err)
	}
	if users.CheckPassword(user, badPassword) {
		t.Error("password should not match")
	}
}
