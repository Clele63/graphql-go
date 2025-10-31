package wrapper

import (
	"context"
	"workbench/graphql-app/queries/generated"
)

func (w *WrappedQueries) CreateUser(ctx context.Context, arg generated.CreateUserParams) error {
	return w.inner.CreateUser(ctx, arg)
}

func (w *WrappedQueries) DeleteUser(ctx context.Context, id string) error {
	return w.inner.DeleteUser(ctx, id)
}

func (w *WrappedQueries) GetCreatedUser(ctx context.Context) (generated.User, error) {
	return w.inner.GetCreatedUser(ctx)
}

func (w *WrappedQueries) GetUpdatedUser(ctx context.Context, id string) (generated.User, error) {
	return w.inner.GetUpdatedUser(ctx, id)
}

func (w *WrappedQueries) GetUserAuthByName(ctx context.Context, name string) (generated.User, error) {
	w.logger.Printf("Fetching user with Name: %s", name)

	user, err := w.inner.GetUserAuthByName(ctx, name)
	if err != nil {
		w.logger.Printf("Error fetching user %s: %v", name, err)
		return generated.User{}, err
	}

	return generated.User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}, nil
}

func (w *WrappedQueries) GetUserByID(ctx context.Context, id string) (generated.User, error) {
	w.logger.Printf("Fetching user with ID: %s", id)

	user, err := w.inner.GetUserByID(ctx, id)
	if err != nil {
		w.logger.Printf("Error fetching user %s: %v", id, err)
		return generated.User{}, err
	}

	return generated.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		CreationDate: user.CreationDate,
	}, nil
}

func (w *WrappedQueries) ListUsers(ctx context.Context) ([]generated.User, error) {
	w.logger.Printf("Fetching users")

	rows, err := w.inner.ListUsers(ctx)
	if err != nil {
		w.logger.Printf("Error fetching users: %v", err)
		return []generated.User{}, err
	}

	users := make([]generated.User, 0, len(rows))
	for _, r := range rows {
		users = append(users, generated.User{
			ID:           r.ID,
			Name:         r.Name,
			Email:        r.Email,
			CreationDate: r.CreationDate,
		})
	}
	return users, nil
}

func (w *WrappedQueries) SearchUsersByName(ctx context.Context, lower string) ([]generated.User, error) {
	w.logger.Printf("Fetching users")

	rows, err := w.inner.SearchUsersByName(ctx, lower)
	if err != nil {
		w.logger.Printf("Error fetching users: %v", err)
		return []generated.User{}, err
	}

	users := make([]generated.User, 0, len(rows))
	for _, r := range rows {
		users = append(users, generated.User{
			ID:           r.ID,
			Name:         r.Name,
			Email:        r.Email,
			CreationDate: r.CreationDate,
		})
	}
	return users, nil
}

func (w *WrappedQueries) UpdateUser(ctx context.Context, arg generated.UpdateUserParams) error {
	return w.inner.UpdateUser(ctx, arg)
}
