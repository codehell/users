package firestore_test

import (
	"errors"
	"github.com/codehell/users"
	"github.com/codehell/users/application"
	"github.com/codehell/users/infrastructure/repositories/firestore"
	"github.com/codehell/users/tests/shared"
	"testing"
)

func TestStoreUser(t *testing.T) {
	repo, err := firestore.NewRepo("codehell-users")
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := shared.GetTestingUser()
	if err := application.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestUserAlreadyError(t *testing.T) {
	repo, err := firestore.NewRepo("codehell-users")
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func () {
		if err := repo.Close(); err != nil {
			t.Log(err)
		}
	}()
	user := shared.GetTestingUser()
	if err := application.StoreUser(repo, user); err != nil {
		t.Fatalf("can not store user: %v", err)
	}
	err = application.StoreUser(repo, user)
	if !errors.Is(users.UserAlreadyExistsError, err) {
		t.Errorf("expected error %v, got %v", users.UserAlreadyExistsError, err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestAllUsers(t *testing.T) {
	repo, err := firestore.NewRepo("codehell-users")
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	if err := shared.CreateTenUsers(repo); err != nil {
		t.Error(err)
	}
	myUsers, err := application.AllUsers(repo)
	if err != nil {
		t.Fatal(err)
	}

	if len(myUsers) != 10 {
		t.Errorf("got %d users, expect %d users", len(myUsers), 20)
	}

	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestFind(t *testing.T) {
	repo, err := firestore.NewRepo("codehell-users")
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := shared.GetTestingUser()
	if err := application.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}

	user, err = application.Find(repo, user.Id())
	if err != nil {
		t.Fatalf("can not find user: %v", err)
	}

	if user.Email() != "cazaplanetas@gmail.com" {
		t.Errorf("user is not the expected user: %v", user)
	}

	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestFindByField(t *testing.T) {
	repo, err := firestore.NewRepo("codehell-users")
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := shared.GetTestingUser()
	if err := application.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}

	user, err = application.FindByField(repo, user.Email(), "email")
	if err != nil {
		t.Fatal("can not find user")
	}

	if user.Email() != "cazaplanetas@gmail.com" {
		t.Errorf("user is not the expected user: %v", user)
	}

	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}
