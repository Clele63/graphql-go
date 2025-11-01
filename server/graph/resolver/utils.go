package resolver

import (
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/queries/generated"
)

func sqlcUserToGraphUser(u generated.User) *model.User {
	return &model.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		CreationDate: u.CreationDate,
	}
}
