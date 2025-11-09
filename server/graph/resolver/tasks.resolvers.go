package resolver

import (
	"context"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
)

func (r *taskResolver) Status(ctx context.Context, obj *model.Task) (string, error) {
	return obj.Column.Name, nil
}

func (r *taskResolver) Assignees(ctx context.Context, obj *model.Task) ([]model.User, error) {
	rows, err := r.Queries.ListAssigneesByTask(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	out := make([]model.User, 0, len(rows))
	for _, t := range rows {
		out = append(out, *SqlcUserToGraphUser(t))
	}
	return out, nil
}

func (r *taskResolver) Column(ctx context.Context, obj *model.Task) (*model.Column, error) {
	row, err := r.Queries.GetColumnForTask(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return SqlcColumnToGraphColumn(row), nil
}

func (r *taskResolver) Comments(ctx context.Context, obj *model.Task) ([]model.Comment, error) {
	rows, err := r.Queries.ListCommentsByTask(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	out := make([]model.Comment, 0, len(rows))
	for _, t := range rows {
		out = append(out, *SqlcCommentToGraphComment(t))
	}
	return out, nil
}

func (r *taskResolver) CreatedAt(ctx context.Context, obj *model.Task) (*scalar.Date, error) {
	return Now(), nil
}
