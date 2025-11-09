package wrapper

import (
	"context"
	"workbench/graphql-app/queries/generated"
)

// Basic CRUD
func (w *WrappedQueries) CreateComment(ctx context.Context, arg generated.CreateCommentParams) error {
	return w.inner.CreateComment(ctx, arg)
}

func (w *WrappedQueries) GetCreatedComment(ctx context.Context) (generated.Comment, error) {
	return w.inner.GetCreatedComment(ctx)
}

func (w *WrappedQueries) GetComment(ctx context.Context, id string) (generated.Comment, error) {
	return w.inner.GetComment(ctx, id)
}

func (w *WrappedQueries) UpdateComment(ctx context.Context, arg generated.UpdateCommentParams) error {
	return w.inner.UpdateComment(ctx, arg)
}

func (w *WrappedQueries) ListComments(ctx context.Context) ([]generated.Comment, error) {
	return w.inner.ListComments(ctx)
}

func (w *WrappedQueries) DeleteComment(ctx context.Context, id string) error {
	return w.inner.DeleteComment(ctx, id)
}

// Extra
func (w *WrappedQueries) GetCommentForTask(ctx context.Context, taskID string) (generated.Comment, error) {
	return w.inner.GetCommentForTask(ctx, taskID)
}

func (w *WrappedQueries) ListCommentsByTask(ctx context.Context, taskID string) ([]generated.Comment, error) {
	return w.inner.ListCommentsByTask(ctx, taskID)
}

func (w *WrappedQueries) GetCommentAuthor(ctx context.Context, id string) (string, error) {
	return w.inner.GetCommentAuthor(ctx, id)
}
