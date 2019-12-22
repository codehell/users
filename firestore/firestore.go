package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/codehell/users"
	"google.golang.org/api/iterator"
)

const CollectionName = "users"

type UsersClient struct {
	projectID string
	client    *firestore.Client
	ctx context.Context
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
	_, err := uf.client.Collection(CollectionName).Doc(u.Username()).Set(uf.ctx, map[string]interface{}{
		"username": u.Username(),
		"email": u.Email(),
		"password": u.Password(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (uf *UsersClient) GetUserByEmail(email string) (users.User, error) {
	var user users.User
	iter := uf.client.Collection(CollectionName).Where("email", "==", email).Documents(uf.ctx)
	doc, err := iter.Next()
	if err != nil {
		return user, err
	}
	dataToUser(doc, &user)
	return user, nil
}

func (uf *UsersClient) GetAll() ([]users.User, error) {
	user := users.User{}
	var users []users.User
	iter := uf.client.Collection(CollectionName).Documents(uf.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return users, err
		}
		dataToUser(doc, &user)
		users = append(users, user)
	}
	return users, nil
}

func (uf *UsersClient) deleteAll() error {
	ref := uf.client.Collection(CollectionName)
	return deleteCollection(uf.ctx, uf.client, ref, 100)
}

func dataToUser(doc *firestore.DocumentSnapshot, user *users.User) {
	data := doc.Data()
	user.SetUsername(fmt.Sprintf("%v", doc.Ref.ID))
	user.SetEmail(fmt.Sprintf("%v", data["email"]))
	user.SetPassword(fmt.Sprintf("%v", data["password"]))
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
