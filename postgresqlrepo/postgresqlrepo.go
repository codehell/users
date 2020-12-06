package postgresqlrepo

import (
	"log"
	"time"

	"github.com/codehell/users"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

type User struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewRepo(db *sqlx.DB) (*UserRepo, error) {
	return &UserRepo{db}, nil
}

func (r *UserRepo) Close() error {
	return r.db.DB.Close()
}

func (r *UserRepo) Store(user users.User) error {
	insert := `INSERT INTO users (id, username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(insert, user.ID().Value(), user.Username().Value(), user.Email(), user.Password(), user.CreatedAt(), user.UpdatedAt())
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) All() ([]users.User, error) {
	return []users.User{}, nil
}

func (r *UserRepo) Search(id string) (users.User, error) {
	return users.User{}, nil
}

func (r *UserRepo) SearchByField(value string, field string) (users.User, error) {
	repoUser := User{}
	query := "SELECT * FROM users WHERE " + field + " = $1"
	logAndReturn := func(err error) (users.User, error) {
		log.Println(err)
		return users.User{}, err
	}
	if err := r.db.Get(&repoUser, query, value); err != nil {
		return logAndReturn(err)
	}
	userID, err := users.NewUserID(repoUser.ID)
	if err != nil {
		return logAndReturn(err)
	}
	username, err := users.NewUsername(repoUser.Username)
	if err != nil {
		return logAndReturn(err)
	}
	user, err := users.NewUser(userID, username, repoUser.Email, "", repoUser.Username)
	if err != nil {
		return logAndReturn(err)
	}
	return user, nil
}

func (r *UserRepo) DeleteAll() error {
	deleteAll := `TRUNCATE users`
	if _, err := r.db.Exec(deleteAll); err != nil {
		return err
	}
	return nil
}
