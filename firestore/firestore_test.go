package firestore

import (
	"github.com/bxcodec/faker/v3"
	"github.com/codehell/users"
	"testing"
)

func getTestingUser() users.User {
	user := users.User{}
	user.SetUsername("cazaplanetas")
	user.SetEmail("cazaplanetas@gmail.com")
	user.SetPassword("secret")
	return user
}

func createUser() users.User{
	var user users.User
	user.SetUsername(faker.Username())
	user.SetEmail(faker.Email())
	user.SetPassword(faker.Password())
	return user
}

func createTwentyUsers(manager *users.UserManager) error {
	for i := 0; i < 20; i++ {
		user := createUser()
		err := manager.CreateUser(&user)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestCreateUser(t *testing.T) {
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func () {
		err = client.CloseClient()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	manager := users.NewManager(client)
	user := getTestingUser()
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	storedUser, err := manager.GetUserByEmail(user.Email())
	if err != nil {
		t.Fatal(err)
	}
	if storedUser.Username() != user.Username() {
		t.Error("The user was not the expected user")
	}
	if err = client.deleteAll(); err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func() {
		err = client.CloseClient()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	manager := users.NewManager(client)
	if err := createTwentyUsers(manager); err != nil {
		t.Fatal(err)
	}
	myUsers, err := manager.GetUsers()
	if err != nil {
		t.Fatal(err)
	}

	if len(myUsers) != 20 {
		t.Errorf("got %d users, expect %d users", len(myUsers), 20)
	}

	if err = client.deleteAll(); err != nil {
		t.Fatal(err)
	}
}



/*import (
	"github.com/codehell/users"
	"testing"
)

func getTestingUser() users.User {
	user := users.User{}
	user.SetUsername("cazaplanetas")
	user.SetEmail("cazaplanetas@gmail.com")
	user.SetPassword("secret")
	return user
}

func TestCreateUser(t *testing.T) {
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := client.CloseClient(); err != nil {
			t.Fatal(err)
		}
	}()
	manager := users.NewManager(client)
	user := getTestingUser()
	if err = manager.CreateUser(&user); err != nil {
		t.Fatal(err)
	}
	storedUser, err := manager.GetUserByEmail(user.Email())
	if err != nil {
		t.Fatal(err)
	}
	if storedUser.Username() != user.Username() {
		t.Errorf("The user %+v created at %s was not the expected user %+v created at %s",
			storedUser, storedUser.CreatedAt(), user, user.CreatedAt())
	}
	if err = client.deleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := NewClient("codehell-users")
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func() {
		err = client.CloseClient()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	manager := users.NewManager(client)
	user := getTestingUser()
	err = manager.CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	myUsers, err := manager.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
	for _, myUser := range myUsers {
		if myUser.Username() != user.Username() {
			t.Errorf("got %+v, expect %+v", myUser, user)
		}
	}
	if err = client.deleteAll(); err != nil {
		t.Fatal(err)
	}
}
*/