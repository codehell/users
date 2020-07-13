package firestorerepo

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/codehell/users"
	"google.golang.org/api/iterator"
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

func (r *UserRepo) Close() error {
	if err := r.client.Close(); err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Store(user users.User) error {
	localUser := User{
		ID:        user.ID().Value(),
		Username:  user.Username().Value(),
		Email:     user.Email(),
		Password:  user.Password(),
		Role:      user.Role(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
	_, err := r.client.Collection(CollectionName).Doc(user.ID().Value()).Set(r.ctx, localUser)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Find(id string) (users.User, error) {
	var fireUser User
	var user users.User
	doc, err := r.client.Collection(CollectionName).Doc(id).Get(r.ctx)
	if err != nil {
		return user, err
	}
	if err := doc.DataTo(&fireUser); err != nil {
		return user, err
	}
	return dataToUser(fireUser)
}

func (r *UserRepo) FindByField(value string, field string) (users.User, error) {
	iter := r.client.Collection(CollectionName).Where(field, "==", value).Documents(r.ctx)
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

func (r *UserRepo) All() ([]users.User, error) {
	var usersCollection []users.User
	var fireUser User
	iter := r.client.Collection(CollectionName).Documents(r.ctx)
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

func (r *UserRepo) DeleteAll() error {
	ref := r.client.Collection(CollectionName)
	return deleteCollection(r.ctx, r.client, ref, 100)
}

func dataToUser(localUser User) (users.User, error) {
	userName, err := users.NewUsername(localUser.Username)
	if err != nil {
		return users.User{}, err
	}
	userId, err := users.NewUserID(localUser.ID)
	if err != nil {
		return users.User{}, err
	}
	user, err := users.NewUser(userId, userName, localUser.Email, localUser.Password, localUser.Role)
	if err != nil {
		return users.User{}, err
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
