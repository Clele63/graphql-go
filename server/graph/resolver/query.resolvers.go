package resolver

import (
	"context"
	"database/sql"
	"fmt"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/middlewares"
	"workbench/graphql-app/queries/generated"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	claims, ok := ctx.Value(middlewares.AuthCtxKey).(*model.UserClaims)
	if !ok {
		return nil, fmt.Errorf("access denied")
	}

	dbUser, err := r.Queries.GetUser(ctx, claims.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found")
		}
		return nil, err
	}

	return SqlcUserToGraphUser(dbUser), nil
}

func (r *queryResolver) Board(ctx context.Context) (*model.Board, error) {
	dbBoard, err := r.Queries.GetBoard(ctx, "b1")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no board found")
		}
		return nil, err
	}

	return SqlcBoardToGraphBoard(dbBoard), nil
}

func (r *queryResolver) Column(ctx context.Context, id string) (*model.Column, error) {
	row, err := r.Queries.GetColumn(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no column found")
		}
		return nil, err
	}

	return SqlcColumnToGraphColumn(row), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]model.User, error) {
	dbUsers, err := r.Queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for _, dbUser := range dbUsers {
		users = append(users, *SqlcUserToGraphUser(dbUser))
	}
	return users, nil
}

func (r *queryResolver) Search(ctx context.Context, term string) ([]model.SearchResult, error) {
	dbUsers, err := r.Queries.SearchUsers(ctx, generated.SearchUsersParams{
		Name:  fmt.Sprintf("%%%s%%", term),
		Email: fmt.Sprintf("%%%s%%", term),
	})

	if err != nil {
		return nil, err
	}

	var results []model.SearchResult
	for _, dbUser := range dbUsers {
		results = append(results, SqlcUserToGraphUser(dbUser))
	}
	return results, nil
}
