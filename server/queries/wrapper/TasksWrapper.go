package wrapper

import (
	"context"
	"workbench/graphql-app/queries/generated"
)

// Basic CRUD
func (w *WrappedQueries) CreateTask(ctx context.Context, arg generated.CreateTaskParams) error {
	return w.inner.CreateTask(ctx, arg)
}

func (w *WrappedQueries) GetCreatedTask(ctx context.Context) (generated.Task, error) {
	return w.inner.GetCreatedTask(ctx)
}

func (w *WrappedQueries) GetTask(ctx context.Context, id string) (generated.Task, error) {
	return w.inner.GetTask(ctx, id)
}

func (w *WrappedQueries) UpdateTask(ctx context.Context, arg generated.UpdateTaskParams) error {
	return w.inner.UpdateTask(ctx, arg)
}

func (w *WrappedQueries) ListTasks(ctx context.Context) ([]generated.Task, error) {
	return w.inner.ListTasks(ctx)
}

func (w *WrappedQueries) DeleteTask(ctx context.Context, id string) error {
	return w.inner.DeleteTask(ctx, id)
}

// Extra
func (w *WrappedQueries) ListTasksForColumnPaginated(ctx context.Context, arg generated.ListTasksForColumnPaginatedParams) ([]generated.Task, error) {
	return w.inner.ListTasksForColumnPaginated(ctx, arg)
}

func (w *WrappedQueries) GetTaskForComment(ctx context.Context, id string) (generated.Task, error) {
	return w.inner.GetTaskForComment(ctx, id)
}

func (w *WrappedQueries) MoveTask(ctx context.Context, arg generated.MoveTaskParams) error {
	return w.inner.MoveTask(ctx, arg)
}
