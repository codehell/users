package users

type UserRepo interface {
	Close() error
	StoreUser(*User) error
	DeleteAll() error
	GetAll() ([]User, error)
	GetUserByEmail(string) (User, error)
}
