package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/codehell/users"
	"google.golang.org/api/iterator"
	"time"
)

const CollectionName = "users"

type Client struct {
	projectID string
	client    *firestore.Client
	ctx       context.Context
}

type User struct {
	Username  string    `firestore:"username"`
	Email     string    `firestore:"email"`
	Password  string    `firestore:"password"`
	Role      string    `firestore:"role"`
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
}

func NewClient(projectID string) (*Client, error) {
	uf := new(Client)
	uf.projectID = projectID
	uf.ctx = context.Background()
	var err error
	uf.client, err = firestore.NewClient(uf.ctx, projectID)
	if err != nil {
		return nil, err
	}
	return uf, nil
}

func (uf *Client) Close() error {
	err := uf.client.Close()
	if err != nil {
		return err
	}
	return nil
}

func (uf *Client) Create(u *users.User) error {
	user := User{
		Username:  u.Username(),
		Email:     u.Email(),
		Password:  u.Password(),
		Role:      u.GetRole(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	_, err := uf.client.Collection(CollectionName).Doc(u.Username()).Set(uf.ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (uf *Client) GetUserByEmail(email string) (users.User, error) {
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
	user = dataToUser(fireUser)
	return user, nil
}

func (uf *Client) GetAll() ([]users.User, error) {
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
		user := dataToUser(fireUser)
		u = append(u, user)
	}
	return u, nil
}

func (uf *Client) DeleteAll() error {
	ref := uf.client.Collection(CollectionName)
	return deleteCollection(uf.ctx, uf.client, ref, 100)
}

func dataToUser(firUser User) users.User {
	var user users.User
	user.SetUsername(firUser.Username)
	user.SetEmail(firUser.Email)
	user.SetPassword(firUser.Password)
	user.SetRole(firUser.Role)
	user.CreatedAt = firUser.CreatedAt
	user.UpdatedAt = firUser.UpdatedAt
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
