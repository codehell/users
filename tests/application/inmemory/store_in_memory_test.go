package inmemory_test

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
	defer repo.Close()
	err = repo.StoreUser(&user)
	if err != nil {
		t.Fatal(err)
	}
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
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	if err := shared.CreateTwentyUsers(repo); err != nil {
		t.Error(err)
	}
	myUsers, err := application.AllUsers(repo)
	if err != nil {
		t.Fatal(err)
	}

	if len(myUsers) != 20 {
		t.Errorf("got %d users, expect %d users", len(myUsers), 20)
	}

	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}


