package resolver

import (
	"workbench/graphql-app/db"
	"workbench/graphql-app/graph/exec"
)

type Resolver struct {
	Services *db.Services
}

func (r *Resolver) Query() exec.QueryResolver       { return &queryResolver{r} }
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
