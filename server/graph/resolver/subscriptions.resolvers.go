package resolver

import (
	"context"
	"fmt"
	"workbench/graphql-app/graph/model"
)

func (r *subscriptionResolver) TaskCreated(ctx context.Context, boardId string) (<-chan *model.Task, error) {
	topic := fmt.Sprintf("taskCreated:%s", boardId)
	ch := make(chan *model.Task)

	r.SubsMutex.Lock()
	r.TaskSubs[topic] = append(r.TaskSubs[topic], ch)
	r.SubsMutex.Unlock()

	go func() {
		<-ctx.Done()
		r.SubsMutex.Lock()
		subs := r.TaskSubs[topic]
		for i, sub := range subs {
			if sub == ch {
				r.TaskSubs[topic] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
		r.SubsMutex.Unlock()
	}()

	return ch, nil
}

func (r *subscriptionResolver) TaskUpdated(ctx context.Context, boardId string) (<-chan *model.Task, error) {
	topic := fmt.Sprintf("taskUpdated:%s", boardId)
	ch := make(chan *model.Task)

	r.SubsMutex.Lock()
	r.TaskSubs[topic] = append(r.TaskSubs[topic], ch)
	r.SubsMutex.Unlock()

	go func() {
		<-ctx.Done()
		r.SubsMutex.Lock()
		subs := r.TaskSubs[topic]
		for i, sub := range subs {
			if sub == ch {
				r.TaskSubs[topic] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
		r.SubsMutex.Unlock()
	}()

	return ch, nil
}

func (r *subscriptionResolver) TaskMoved(ctx context.Context, boardId string) (<-chan *model.Task, error) {
	topic := fmt.Sprintf("taskMoved:%s", boardId)
	ch := make(chan *model.Task)

	r.SubsMutex.Lock()
	r.TaskSubs[topic] = append(r.TaskSubs[topic], ch)
	r.SubsMutex.Unlock()

	go func() {
		<-ctx.Done()
		r.SubsMutex.Lock()
		subs := r.TaskSubs[topic]
		for i, sub := range subs {
			if sub == ch {
				r.TaskSubs[topic] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
		r.SubsMutex.Unlock()
	}()

	return ch, nil
}

func (r *subscriptionResolver) CommentAdded(ctx context.Context, taskID string) (<-chan *model.Comment, error) {
	topic := fmt.Sprintf("commentAdded:%s", taskID)
	ch := make(chan *model.Comment)

	r.SubsMutex.Lock()
	r.CommentSubs[topic] = append(r.CommentSubs[topic], ch)
	r.SubsMutex.Unlock()

	go func() {
		<-ctx.Done()
		r.SubsMutex.Lock()
		subs := r.CommentSubs[topic]
		for i, sub := range subs {
			if sub == ch {
				r.CommentSubs[topic] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
		r.SubsMutex.Unlock()
	}()

	return ch, nil
}
