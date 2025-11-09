package resolver

import (
	"context"
	"sort"
	"workbench/graphql-app/graph/model"
)

// Columns is the resolver for the columns field.
func (r *boardResolver) Columns(ctx context.Context, obj *model.Board) ([]model.Column, error) {
	rows, err := r.Queries.ListColumnsByBoard(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	sort.Slice(rows, func(a, b int) bool {
		return rows[a].Order < rows[b].Order
	})

	out := make([]model.Column, 0, len(rows))
	for _, t := range rows {
		out = append(out, *SqlcColumnToGraphColumn(t))
	}
	return out, nil
}
