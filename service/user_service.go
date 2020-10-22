package service

import (
	"context"
	"playing-with-golang-on-k8s/api"
	"playing-with-golang-on-k8s/crypto"
	"playing-with-golang-on-k8s/store"
	"strings"
	"time"
)

//UserService service logic for users
type UserService struct {
	userStore store.UserStore
}

//NewUserService constructs a new UserService
func NewUserService(userStore store.UserStore) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

//Create user creation logic
func (us *UserService) Create(ctx context.Context, req *api.UserCreationRequest, createdAt time.Time) (*store.User, error) {
	email := strings.ToLower(req.Email)
	taken, err := us.userStore.IsEmailTaken(ctx, email)

	if err != nil {
		return nil, err
	}
	if taken {
		return nil, store.ErrEmailTaken
	}

	//Encrypt the password
	password, err := crypto.Encrypt(req.Password)
	if err != nil {
		return nil, err
	}

	user := req.ToStore(createdAt)
	user.Password = password
	user.UpdatedAt = createdAt
	user.Email = email
	createdUser, err := us.userStore.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

//Get returns a user based on its unique ID
func (us *UserService) Get(ctx context.Context, id string) (*store.User, error) {
	return us.userStore.GetUser(ctx, id)
}
