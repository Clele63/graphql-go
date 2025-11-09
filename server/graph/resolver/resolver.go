package resolver

import (
	"sync"
	"workbench/graphql-app/graph/exec"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/queries/wrapper"
)

type Resolver struct {
	Queries *wrapper.WrappedQueries

	TaskSubs    map[string][]chan *model.Task
	CommentSubs map[string][]chan *model.Comment

	SubsMutex sync.Mutex
}

func (r *Resolver) Query() exec.QueryResolver               { return &queryResolver{r} }
func (r *Resolver) Mutation() exec.MutationResolver         { return &mutationResolver{r} }
func (r *Resolver) Subscription() exec.SubscriptionResolver { return &subscriptionResolver{r} }
func (r *Resolver) Board() exec.BoardResolver               { return &boardResolver{r} }
func (r *Resolver) Column() exec.ColumnResolver             { return &columnResolver{r} }
func (r *Resolver) Comment() exec.CommentResolver           { return &commentResolver{r} }
func (r *Resolver) Task() exec.TaskResolver                 { return &taskResolver{r} }
func (r *Resolver) User() exec.UserResolver                 { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type boardResolver struct{ *Resolver }
type columnResolver struct{ *Resolver }
type commentResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
