package wrapper

import (
	"context"
	"workbench/graphql-app/queries/generated"
)

// Basic CRUD
func (w *WrappedQueries) CreateBoard(ctx context.Context, arg generated.CreateBoardParams) error {
	return w.inner.CreateBoard(ctx, arg)
}

func (w *WrappedQueries) GetCreatedBoard(ctx context.Context) (generated.Board, error) {
	return w.inner.GetCreatedBoard(ctx)
}

func (w *WrappedQueries) GetBoard(ctx context.Context, id string) (generated.Board, error) {
	return w.inner.GetBoard(ctx, id)
}

func (w *WrappedQueries) UpdateBoard(ctx context.Context, arg generated.UpdateBoardParams) error {
	return w.inner.UpdateBoard(ctx, arg)
}

func (w *WrappedQueries) ListBoards(ctx context.Context) ([]generated.Board, error) {
	return w.inner.ListBoards(ctx)
}

func (w *WrappedQueries) DeleteBoard(ctx context.Context, id string) error {
	return w.inner.DeleteBoard(ctx, id)
}

// Extra
func (w *WrappedQueries) GetBoardForColumn(ctx context.Context, id string) (generated.Board, error) {
	return w.inner.GetBoardForColumn(ctx, id)
}
