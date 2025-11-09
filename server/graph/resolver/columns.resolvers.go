package resolver

import (
	"context"
	"encoding/base64"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/queries/generated"
)

// Order is the resolver for the order field.
func (r *columnResolver) Order(ctx context.Context, obj *model.Column) (int32, error) {
	return int32(obj.Order), nil
}

// Tasks is the resolver for the paginated tasks field.
func (r *columnResolver) Tasks(ctx context.Context, obj *model.Column, first *int32, after *string) (*model.TaskConnection, error) {
	limit := int32(10)
	if first != nil {
		limit = *first
	}
	queryLimit := int32(limit + 1)

	cursorID := ""
	if after != nil {
		cursorID = decodeCursor(*after)
	}

	rows, err := r.Queries.ListTasksForColumnPaginated(ctx, generated.ListTasksForColumnPaginatedParams{
		ColumnID: obj.ID,
		ID:       cursorID,
		Limit:    queryLimit,
	})
	if err != nil {
		return nil, err
	}

	edges := []*model.TaskEdge{}
	hasNextPage := len(rows) == int(queryLimit)
	var endCursor *string

	nodesToProcess := rows
	if hasNextPage {
		nodesToProcess = rows[:limit]
	}

	for _, dbTask := range nodesToProcess {
		gqlTask := SqlcTaskToGraphTask(dbTask)

		cursor := encodeCursor(dbTask.ID)

		edges = append(edges, &model.TaskEdge{
			Node:   gqlTask,
			Cursor: cursor,
		})
	}

	if len(edges) > 0 {
		endCursor = &edges[len(edges)-1].Cursor
	}

	return &model.TaskConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
	}, nil
}

func encodeCursor(id string) string {
	return base64.StdEncoding.EncodeToString([]byte(id))
}

func decodeCursor(cur string) string {
	b, _ := base64.StdEncoding.DecodeString(cur)
	return string(b)
}
