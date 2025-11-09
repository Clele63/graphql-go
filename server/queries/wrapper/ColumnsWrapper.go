package wrapper

import (
	"context"
	"workbench/graphql-app/queries/generated"
)

// Basic CRUD
func (w *WrappedQueries) CreateColumn(ctx context.Context, arg generated.CreateColumnParams) error {
	return w.inner.CreateColumn(ctx, arg)
}

func (w *WrappedQueries) GetCreatedColumn(ctx context.Context) (generated.Column, error) {
	return w.inner.GetCreatedColumn(ctx)
}

func (w *WrappedQueries) GetColumn(ctx context.Context, id string) (generated.Column, error) {
	return w.inner.GetColumn(ctx, id)
}

func (w *WrappedQueries) UpdateColumn(ctx context.Context, arg generated.UpdateColumnParams) error {
	return w.inner.UpdateColumn(ctx, arg)
}

func (w *WrappedQueries) ListColumns(ctx context.Context) ([]generated.Column, error) {
	return w.inner.ListColumns(ctx)
}

func (w *WrappedQueries) DeleteColumn(ctx context.Context, id string) error {
	return w.inner.DeleteColumn(ctx, id)
}

// Extra
func (w *WrappedQueries) ListColumnsByBoard(ctx context.Context, boardID string) ([]generated.Column, error) {
	return w.inner.ListColumnsByBoard(ctx, boardID)
}

func (w *WrappedQueries) GetColumnForTask(ctx context.Context, id string) (generated.Column, error) {
	return w.inner.GetColumnForTask(ctx, id)
}
