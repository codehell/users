package postgresqlrepo_test

import (
	"errors"
	"testing"

	"github.com/codehell/users"
	"github.com/codehell/users/postgresqlrepo"
	"github.com/codehell/users/testsutils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestStoreUser(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=postgres password=secret dbname=users sslmode=disable")
	if err != nil {
		t.Fatalf("can not connect to database: %v", err)
	}
	repo, err := postgresqlrepo.NewRepo(db)
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
	db, err := sqlx.Connect("postgres", "user=postgres password=secret dbname=users sslmode=disable")
	if err != nil {
		t.Fatalf("can not connect to database: %v", err)
	}
	repo, err := postgresqlrepo.NewRepo(db)
	if err != nil {
		t.Fatalf("can not create repo %v", err)
	}
	defer repo.Close()
	user := testsutils.GetTestingUser()
	if err := users.StoreUser(repo, user); err != nil {
		t.Errorf("can not store user: %v", err)
	}
	user = testsutils.GetTestingUser()
	userID, _ := users.NewUserID(user.ID().Value())
	userName, _ := users.NewUsername(user.Username().Value())
	user, _ = users.NewUser(userID, userName, "cazaplanetas@hotmail.com", "secret", "normal")
	err = users.StoreUser(repo, user)
	if !errors.Is(&users.UserAlreadyExistsError, err) {
		t.Errorf("expected error %v, got %v", users.UserAlreadyExistsError, err)
	}
	if err := repo.DeleteAll(); err != nil {
		t.Fatalf("can not delete table")
	}
}
