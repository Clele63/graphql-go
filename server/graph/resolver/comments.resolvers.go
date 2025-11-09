package resolver

import (
	"context"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
)

// Author is the resolver for the author field
func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (*model.User, error) {
	row, err := r.Queries.GetUserForComment(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return SqlcUserToGraphUser(row), nil
}

// Task is the resolver for the task field
func (r *commentResolver) Task(ctx context.Context, obj *model.Comment) (*model.Task, error) {
	row, err := r.Queries.GetTaskForComment(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return SqlcTaskToGraphTask(row), nil
}

// CreatedAt is the resolver for the createdAt field
func (r *commentResolver) CreatedAt(ctx context.Context, obj *model.Comment) (*scalar.Date, error) {
	return Now(), nil
}
