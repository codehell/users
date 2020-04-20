package application_test

import (
	"github.com/codehell/users"
	"github.com/codehell/users/application"
	"github.com/codehell/users/infrastructure/repositories/inmemory"
	"github.com/codehell/users/tests/shared"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := shared.GetTestingUser()
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal(err)
	}
	err = repo.StoreUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if repo != nil {
			if err := repo.Close(); err != nil {
				t.Fatal(err)
			}
		}
	}()
	storedUser, err := repo.GetUserByEmail(user.Email())
	if err != nil {
		t.Fatal(err)
	}
	if storedUser.Email() != user.Email() {
		t.Error("The user was not the expected user")
	}
	if err = repo.DeleteAll(); err != nil {
		t.Error(err)
	}
}

func TestStoreUser(t *testing.T) {
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			t.Log("can not close repo")
		}
	}()
	user := shared.GetTestingUser()
	if err := application.StoreUser(user, repo); err != nil {
		t.Error(err)
	}
	if err := application.StoreUser(user, repo); err != users.ErrUserAlreadyExists {
		t.Error(err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestGetUsers(t *testing.T) {
	client, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create client")
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Fatal("can not close client")
		}
	}()
	if err := shared.CreateTwentyUsers(client); err != nil {
		t.Error(err)
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
	user := shared.GetTestingUser()
	client, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal(err)
	}
	err = client.StoreUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = client.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Error(err)
	}
	if !users.CheckPassword(user.Password(), "secret1") {
		t.Error("password should match")
	}
}

func TestWrongPassword(t *testing.T) {
	user := shared.GetTestingUser()
	client, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal(err)
	}
	err = client.StoreUser(&user)
	if err != nil {
		t.Fatal(err)
	}
	badPassword := "badPassword"
	if users.CheckPassword(user.Password(), badPassword) {
		t.Error("password should not match")
	}
}
