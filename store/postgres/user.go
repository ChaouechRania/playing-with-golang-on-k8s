package postgres

import (
	"playing-with-golang-on-k8s/crypto"
	"playing-with-golang-on-k8s/store"
	"context"
	"strings"
	"github.com/jinzhu/gorm"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/pkg/errors"
)

//Authenticate returns a user based on its unique ID
func (c *Client) Authenticate(ctx context.Context, email string, password string) (*store.User, error) {
	user := new(store.User)
	err := c.db.Where(store.User{Email: email}).Take(user).Error
	if err != nil {
		if err == store.ErrNoSuchEntity {
			return nil, store.ErrNoSuchEntity
		}
		return nil, err
	}
	err = crypto.Compare(user.Password, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//IsEmailTaken tests if the provided email address is already taken
func (c *Client) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	count := 0
	r := c.db.Model(&store.User{}).Where("email = ?", strings.ToLower(email)).
		Count(&count)
	err := r.Error
	if err != nil {
		return false, err
	}

	return count > 0, err
}

//GetUser returns a user based on its unique ID
func (c *Client) GetUser(ctx context.Context, id string) (*store.User, error) {
	user := new(store.User)
	err := c.db.Where(store.User{ID: id}).Take(user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, store.ErrNoSuchEntity
		}
		return nil, err
	}

	return user, nil
}

//CreateUser creates a user
func (c *Client) CreateUser(ctx context.Context, user *store.User) (*store.User, error) {
	user.Email = strings.ToLower(user.Email)
	taken, err := c.IsEmailTaken(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if taken {
		return nil, store.ErrEmailTaken
	}

	id, err := gonanoid.Nanoid()
	if err != nil {
		return nil, errors.Wrap(err, "generating job id")
	}
	user.ID = id
	err = c.db.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
