package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/codehell/users"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

const CollectionName = "users"

type UsersClient struct {
	projectID string
	client    *firestore.Client
	ctx context.Context
}

type User struct {
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(projectID string) (*UsersClient, error) {
	uf := new(UsersClient)
	uf.projectID = projectID
	uf.ctx = context.Background()
	var err error
	uf.client, err = firestore.NewClient(uf.ctx, projectID)
	if err != nil {
		return uf, err
	}
	return uf, nil
}

func (uf *UsersClient) CloseClient() error {
	err := uf.client.Close()
	if err != nil {
		return err
	}
	return nil
}

func (uf *UsersClient) Create(u *users.User) error {
	user := User{
		Username:  u.Username(),
		Email:     u.Email(),
		Password:  u.Password(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}
	res, err := uf.client.Collection(CollectionName).Doc(u.Username()).Set(uf.ctx, user)
	log.Println("wtf-time", res.UpdateTime.String(), user)
	if err != nil {
		return err
	}
	return nil
}

func (uf *UsersClient) GetUserByEmail(email string) (users.User, error) {
	var fireUser User
	var user users.User
	iter := uf.client.Collection(CollectionName).Where("email", "==", email).Documents(uf.ctx)
	doc, err := iter.Next()
	if err != nil {
		return user, err
	}
	if err := doc.DataTo(&fireUser); err != nil {
		return user, err
	}
	user = dataToUser(doc.Ref.ID, fireUser)
	return user, nil
}

func (uf *UsersClient) GetAll() ([]users.User, error) {
	var u []users.User
	var fireUser User
	iter := uf.client.Collection(CollectionName).Documents(uf.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return u, err
		}
		user := dataToUser(doc.Ref.ID, fireUser)
		u = append(u, user)
	}
	return u, nil
}

func (uf *UsersClient) deleteAll() error {
	ref := uf.client.Collection(CollectionName)
	return deleteCollection(uf.ctx, uf.client, ref, 100)
}

func dataToUser(ID string, firUser User) users.User {
	var user users.User
	user.SetUsername(ID)
	user.SetEmail(firUser.Email)
	user.SetPassword(firUser.Password)
	user.SetCreatedAt(firUser.CreatedAt)
	user.SetUpdatedAt(firUser.UpdatedAt)
	return user
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
