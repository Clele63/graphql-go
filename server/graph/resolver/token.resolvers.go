package resolver

import (
	"context"
	"workbench/graphql-app/graph/model"
)

// ExpiredAt is the resolver for the expired_at field.
func (r *tokenResolver) ExpiredAt(ctx context.Context, obj *model.Token) (int32, error) {
	return int32(obj.ExpiredAt.Unix()), nil
}
