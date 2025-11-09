package wrapper

import (
	"context"
	"workbench/graphql-app/queries/generated"
)

// Basic CRUD
func (w *WrappedQueries) CreateUser(ctx context.Context, arg generated.CreateUserParams) error {
	return w.inner.CreateUser(ctx, arg)
}

func (w *WrappedQueries) GetCreatedUser(ctx context.Context) (generated.User, error) {
	return w.inner.GetCreatedUser(ctx)
}

func (w *WrappedQueries) GetUser(ctx context.Context, id string) (generated.User, error) {
	return w.inner.GetUser(ctx, id)
}

func (w *WrappedQueries) UpdateUser(ctx context.Context, arg generated.UpdateUserParams) error {
	return w.inner.UpdateUser(ctx, arg)
}

func (w *WrappedQueries) ListUsers(ctx context.Context) ([]generated.User, error) {
	return w.inner.ListUsers(ctx)
}

func (w *WrappedQueries) DeleteUser(ctx context.Context, id string) error {
	return w.inner.DeleteUser(ctx, id)
}

// Extra
func (w *WrappedQueries) SearchUsers(ctx context.Context, arg generated.SearchUsersParams) ([]generated.User, error) {
	return w.inner.SearchUsers(ctx, arg)
}

func (w *WrappedQueries) GetUserForComment(ctx context.Context, id string) (generated.User, error) {
	return w.inner.GetUserForComment(ctx, id)
}

func (w *WrappedQueries) GetUserByEmail(ctx context.Context, email string) (generated.User, error) {
	return w.inner.GetUserByEmail(ctx, email)
}
