package resolver

import (
	"context"
	"errors"
	"time"

	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
	"workbench/graphql-app/queries/generated"
	"workbench/graphql-app/utils"
)

// ---------------- USERS ----------------

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	err := r.Queries.CreateUser(ctx, generated.CreateUserParams{
		Name:         input.Name,
		Password:     input.Password,
		Email:        input.Email,
		CreationDate: input.CreationDate,
	})

	if err != nil {
		return nil, err
	}

	user, err := r.Queries.GetCreatedUser(ctx)
	if err != nil {
		return nil, err
	}

	return sqlcUserToGraphUser(user), nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	err := r.Queries.UpdateUser(ctx, generated.UpdateUserParams{
		ID:    input.ID,
		Name:  *input.Name,
		Email: *input.Email,
	})
	if err != nil {
		return nil, err
	}

	user, err := r.Queries.GetUpdatedUser(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return sqlcUserToGraphUser(user), nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	err := r.Queries.DeleteUser(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ---------------- LOGIN ----------------

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, name string, password string) (*model.Token, error) {
	user, err := r.Queries.GetUserAuthByName(ctx, name)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !utils.ComparePassword(password, user.Password) {
		return nil, errors.New("passwords doesn't match")
	}

	expiredAt := time.Now().Add(time.Hour * 1)
	obj := &model.Token{
		Token:     utils.GenerateJwt(user.ID, int64(expiredAt.Unix())),
		ExpiredAt: scalar.Date{Time: &expiredAt},
	}

	return obj, nil
}
