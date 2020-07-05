package application_test

import (
	"github.com/codehell/users"
	"github.com/codehell/users/application"
	"github.com/codehell/users/infrastructure/repositories/inmemory"
	tests2 "github.com/codehell/users/tests"
	"reflect"
	"testing"
)

func TestSignIn(t *testing.T) {
	repo, err := inmemory.NewRepo()
	if err != nil {
		t.Fatal("can not create repo")
	}
	defer func() {
		if err := repo.Close(); err != nil {
			t.Fatal("can not close the repo")
		}
	}()
	user := tests2.GetTestingUser()
	if err := application.StoreUser(repo, user); err != nil {
		t.Fatal(err)
	}
	type args struct {
		userRepo users.UserRepo
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    users.User
		wantErr bool
	}{
		{
			"user can signing",
			args{
				userRepo: repo,
				email:    "cazaplanetas@gmail.com",
				password: "secret1",
			},
			user,
			false,
		},
		{
			"user can not signing with wrong password",
			args{
				userRepo: repo,
				email:    "cazaplanetas@gmail.com",
				password: "secret2",
			},
			users.User{},
			true,
		},
		{
			"user can not signing if not exist",
			args{
				userRepo: repo,
				email:    "codhell@gmail.com",
				password: "secret1",
			},
			users.User{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := application.SignIn(tt.args.userRepo, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignIn() got = %v, want %v", got, tt.want)
			}
		})
		if err = repo.DeleteAll(); err != nil {
			t.Error(err)
		}
	}
}
