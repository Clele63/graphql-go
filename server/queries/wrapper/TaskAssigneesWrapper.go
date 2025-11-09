package wrapper

import (
	"context"
	"workbench/graphql-app/queries/generated"
)

func (w *WrappedQueries) AddTaskAssignee(ctx context.Context, arg generated.AddTaskAssigneeParams) error {
	return w.inner.AddTaskAssignee(ctx, arg)
}

func (w *WrappedQueries) GetTaskAssignee(ctx context.Context, arg generated.GetTaskAssigneeParams) (generated.TaskAssignee, error) {
	return w.inner.GetTaskAssignee(ctx, arg)
}

func (w *WrappedQueries) ListAssigneesByTask(ctx context.Context, taskID string) ([]generated.User, error) {
	return w.inner.ListAssigneesByTask(ctx, taskID)
}

func (w *WrappedQueries) ListTaskAssignsByTask(ctx context.Context, taskID string) ([]generated.TaskAssignee, error) {
	return w.inner.ListTaskAssignsByTask(ctx, taskID)
}

func (w *WrappedQueries) ListTaskAssignsByUser(ctx context.Context, userID string) ([]generated.TaskAssignee, error) {
	return w.inner.ListTaskAssignsByUser(ctx, userID)
}

func (w *WrappedQueries) DeleteTaskAssignee(ctx context.Context, arg generated.DeleteTaskAssigneeParams) error {
	return w.inner.DeleteTaskAssignee(ctx, arg)
}

func (w *WrappedQueries) ClearTaskAssigneesByTask(ctx context.Context, taskID string) error {
	return w.inner.ClearTaskAssigneesByTask(ctx, taskID)
}

func (w *WrappedQueries) ClearTaskAssigneesByUser(ctx context.Context, userID string) error {
	return w.inner.ClearTaskAssigneesByUser(ctx, userID)
}
