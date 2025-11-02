package resolver

import (
	"context"
	"errors"
	"time"

	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
	"workbench/graphql-app/middlewares"
	"workbench/graphql-app/utils"
)

// ---------------- USERS ----------------

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	userAuth := middlewares.GetAuthFromContext(ctx)
	if userAuth.UserID == "0" {
		return nil, errors.New("access denied")
	}

	userMap, err := r.Services.Users.CreateUser(input.Name, input.Password, input.Email)
	if err != nil {
		return nil, err
	}

	return badgerUserToGraphUser(userMap), nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	userAuth := middlewares.GetAuthFromContext(ctx)
	if userAuth.UserID == "0" {
		return nil, errors.New("access denied")
	}

	userMap, err := r.Services.Users.UpdateUser(input.ID, input.Name)
	if err != nil {
		return nil, err
	}

	return badgerUserToGraphUser(userMap), nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	userAuth := middlewares.GetAuthFromContext(ctx)
	if userAuth.UserID == "0" {
		return false, errors.New("access denied")
	}

	err := r.Services.Users.DeleteUser(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ---------------- LOGIN ----------------

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, name string, password string) (*model.Token, error) {
	user, err := r.Services.Users.GetUserAuthByName(name)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !utils.ComparePassword(password, user["password"]) {
		return nil, errors.New("passwords doesn't match")
	}

	expiredAt := time.Now().Add(time.Hour * 1)
	userID := user["id"]

	obj := &model.Token{
		Token:     utils.GenerateJwt(userID, int64(expiredAt.Unix())),
		ExpiredAt: scalar.Date{Time: &expiredAt},
	}

	return obj, nil
}
