package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/codehell/users"
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

func NewRepo(projectID string) (*UserRepo, error) {
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

func (uf *UserRepo) Store(u users.User) error {
	user := User{
		ID: u.Id().Value(),
		Username:  u.Username().Value(),
		Email:     u.Email(),
		Password:  u.Password(),
		Role:      u.Role(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}
	_, err := uf.client.Collection(CollectionName).Doc(u.Id().Value()).Set(uf.ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (uf *UserRepo) Find(id string) (users.User, error) {
	var fireUser User
	var user users.User
	doc, err := uf.client.Collection(CollectionName).Doc(id).Get(uf.ctx)
	if err != nil {
		return user, err
	}
	if err := doc.DataTo(&fireUser); err != nil {
		return user, err
	}
	return dataToUser(fireUser)
}

func (uf *UserRepo) FindByField(value string, field string) (users.User, error) {
	iter := uf.client.Collection(CollectionName).Where(field, "==", value).Documents(uf.ctx)
	doc, err := iter.Next()
	if err != nil {
		return users.User{}, err
	}
	var fireUser User
	if err := doc.DataTo(&fireUser); err != nil {
		return users.User{}, err
	}
	return dataToUser(fireUser)
}

func (uf *UserRepo) All() ([]users.User, error) {
	var usersCollection []users.User
	var fireUser User
	iter := uf.client.Collection(CollectionName).Documents(uf.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return usersCollection, err
		}
		if err := doc.DataTo(&fireUser); err != nil {
			return usersCollection, err
		}
		user, err := dataToUser(fireUser)
		if err != nil {
			return usersCollection, err
		}
		usersCollection = append(usersCollection, user)
	}
	return usersCollection, nil
}

func (uf *UserRepo) DeleteAll() error {
	ref := uf.client.Collection(CollectionName)
	return deleteCollection(uf.ctx, uf.client, ref, 100)
}

func dataToUser(fu User) (users.User, error) {
	userName, err := users.NewUsername(fu.Username)
	if err != nil {
		return users.User{}, users.UserSystemError
	}
	userId, err := users.NewUserId(fu.ID)
	if err != nil {
		return users.User{}, users.UserSystemError
	}
	user, err := users.NewUser(userId, userName, fu.Email, fu.Password, fu.Role)
	if err != nil {
		return users.User{}, users.UserSystemError
	}

	return user, nil
}

func deleteCollection(ctx context.Context, client *firestore.Client, ref *firestore.CollectionRef, batchSize int) error {
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
