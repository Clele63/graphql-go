package resolver

import (
	"context"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
)

func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (string, error) {
	return "https://api.dicebear.com/7.x/bottts/svg?seed=" + obj.ID, nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (*scalar.Date, error) {
	return Now(), nil
}
