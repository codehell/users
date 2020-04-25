package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/codehell/users"
	"github.com/codehell/users/valueobjects"
	"google.golang.org/api/iterator"
	"time"
)

const CollectionName = "users"

type UserRepo struct {
	projectID string
	client    *firestore.Client
	ctx       context.Context
}

type User struct {
	ID        string    `firestore:"userId"`
	Username  string    `firestore:"username"`
	Email     string    `firestore:"email"`
	Password  string    `firestore:"password"`
	Role      string    `firestore:"role"`
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
}

func NewClient(projectID string) (*UserRepo, error) {
	uf := new(UserRepo)
	uf.projectID = projectID
	uf.ctx = context.Background()
	var err error
	if uf.client, err = firestore.NewClient(uf.ctx, projectID); err != nil {
		return nil, err
	}
	return uf, nil
}

func (uf *UserRepo) Close() error {
	if err := uf.client.Close(); err != nil {
		return err
	}
	return nil
}

func (uf *UserRepo) StoreUser(u *users.User) error {
	var err error
	user := User{
		Username:  u.Username().Value(),
		Email:     u.Email(),
		Password:  u.Password(),
		Role:      u.Role(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}
	_, err = uf.client.Collection(CollectionName).Doc(u.Email()).Set(uf.ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (uf *UserRepo) GetUserByEmail(email string) (users.User, error) {
	var fireUser User
	var user users.User
	doc, err := uf.client.Collection(CollectionName).Doc(email).Get(uf.ctx)
	if err != nil {
		return user, err
	}
	if err := doc.DataTo(&fireUser); err != nil {
		return user, err
	}
	user, err = dataToUser(fireUser)

	return user, nil
}

func (uf *UserRepo) GetAll() ([]users.User, error) {
	var u []users.User
	var fireUser User
	iter := uf.client.Collection(CollectionName).Documents(uf.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return u, err
		}
		if err := doc.DataTo(&fireUser); err != nil {
			return u, err
		}
		user, err := dataToUser(fireUser)
		if err != nil {
			return u, err
		}
		u = append(u, user)
	}
	return u, nil
}

func (uf *UserRepo) DeleteAll() error {
	ref := uf.client.Collection(CollectionName)
	return deleteCollection(uf.ctx, uf.client, ref, 100)
}

func dataToUser(fu User) (users.User, error) {
	userName, err := valueobjects.NewUsername(fu.Username)
	if err != nil {
		return users.User{}, err
	}
	user, err := users.NewUser(fu.ID, userName, fu.Email, fu.Password, fu.Role)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}

func deleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}