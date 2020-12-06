package firestorerepo_test

import (
	"errors"
	"testing"

	"github.com/codehell/users/firestorerepo"

	"github.com/codehell/users"
	"github.com/codehell/users/testsutils"
)

const projectID = "chlack-20"

func TestStoreUser(t *testing.T) {
	repo, err := firestorerepo.NewRepo(projectID)
	if err != nil {
		t.Fatalf("can not create repo: %v", err)
	}
	defer repo.Close()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestUserAlreadyExistError(t *testing.T) {
	repo, err := firestorerepo.NewRepo(projectID)
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		if err := repo.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Fatalf("can not store user: %v", err)
	}
	err = users.StoreUser(repo, user)
	if !errors.Is(&users.UserAlreadyExistsError, err) {
		t.Errorf("expected error %v, got %v", users.UserAlreadyExistsError, err)
	}
	if err = repo.DeleteAll(); err != nil {
		t.Fatal(err)
	}
}

func TestAllUsers(t *testing.T) {
	repo, err := firestorerepo.NewRepo(projectID)
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	if err := testsutils.CreateTenUsers(repo); err != nil {
		t.Error(err)
	}
	myUsers, err := users.AllUsers(repo)
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
	repo, err := firestorerepo.NewRepo(projectID)
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}

	user, err = users.Search(repo, user.ID())
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
	repo, err := firestorerepo.NewRepo(projectID)
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer repo.Close()
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
