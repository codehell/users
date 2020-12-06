package inmemoryrepo_test

import (
	"errors"
	"log"
	"testing"

	"github.com/codehell/users"
	"github.com/codehell/users/inmemoryrepo"
	"github.com/codehell/users/testsutils"
)

func TestStoreUser(t *testing.T) {
	repo, err := inmemoryrepo.New()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Fatal(err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestUserAlreadyExistError(t *testing.T) {
	repo, err := inmemoryrepo.New()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Print(err)
		}
	}()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Fatal(err)
	}

	if err := users.StoreUser(repo, user); !errors.Is(err, &users.UserAlreadyExistsError) {
		t.Errorf("expected error %v, got %v", users.UserAlreadyExistsError, err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestAllUsers(t *testing.T) {
	repo, err := inmemoryrepo.New()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err := testsutils.CreateTenUsers(repo); err != nil {
		t.Fatal(err)
	}
	myUsers, err := users.AllUsers(repo)
	if err != nil {
		t.Fatal(err)
	}

	if len(myUsers) != 10 {
		t.Errorf("got %d users, expect %d users", len(myUsers), 10)
	}

	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestSearch(t *testing.T) {
	repo, err := inmemoryrepo.New()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Fatal(err)
	}

	user, err = users.Search(repo, user.ID())
	if err != nil {
		t.Fatal("can not find user")
	}

	if user.Email() != "cazaplanetas@gmail.com" {
		t.Error("user is not the expected")
	}
}

func TestFindField(t *testing.T) {
	repo, err := inmemoryrepo.New()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}

	user, err = users.SearchByField(repo, user.Email(), "email")
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
	repo, err := inmemoryrepo.New()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}
	_, err = users.SignIn(repo, "cazaplanetas@gmail.com", "secret1")
	if err != nil {
		t.Error(err)
	}
}
