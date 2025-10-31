package resolver

import (
	"workbench/graphql-app/graph/exec"
	"workbench/graphql-app/queries/wrapper"
)

type Resolver struct {
	Queries *wrapper.WrappedQueries
}

func (r *Resolver) Query() exec.QueryResolver       { return &queryResolver{r} }
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Token() exec.TokenResolver       { return &tokenResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type tokenResolver struct{ *Resolver }
