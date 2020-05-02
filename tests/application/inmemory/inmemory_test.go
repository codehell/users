package inmemory_test

import (
	"github.com/codehell/users/application"
	"github.com/codehell/users/infrastructure/repositories/inmemory"
	"github.com/codehell/users/tests/shared"
	"testing"
)

func TestStoreUser(t *testing.T) {
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := shared.GetTestingUser()
	if err := application.StoreUser(user, repo); err != nil {
		t.Fatal(err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestAllUsers(t *testing.T) {
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	if err := shared.CreateTwentyUsers(repo); err != nil {
		t.Fatal(err)
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

func TestFind(t *testing.T) {
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := shared.GetTestingUser()
	if err := application.StoreUser(user, repo); err != nil {
		t.Fatal(err)
	}

	user, err = application.Find(repo, user.Id())
	if err != nil {
		t.Fatal("can not find user")
	}

	if user.Email() != "cazaplanetas@gmail.com" {
		t.Error("user is not the expected")
	}
}

func TestFindField(t *testing.T) {
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := shared.GetTestingUser()
	if err := application.StoreUser(user, repo); err != nil {
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

func TestSignIn(t *testing.T) {
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := shared.GetTestingUser()
	if err := application.StoreUser(user, repo); err != nil {
		t.Errorf("can not store user: %v", err)
	}
	_, err = application.SignIn(repo, "cazaplanetas@gmail.com", "secret1")
	if err != nil {
		t.Error(err)
	}
}
