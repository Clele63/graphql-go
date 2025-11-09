package model

import "workbench/graphql-app/graph/resolver/scalar"

type Task struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description *string     `json:"description,omitempty"`
	Status      string      `json:"status"`
	Assignees   []*User     `json:"assignees"`
	Column      *Column     `json:"column"`
	Comments    []*Comment  `json:"comments"`
	CreatedAt   scalar.Date `json:"createdAt"`
}

func (Task) IsNode()         {}
func (t Task) GetID() string { return t.ID }

type TaskEdge struct {
	Cursor string `json:"cursor"`
	Node   *Task  `json:"node"`
}

type TaskConnection struct {
	Edges    []*TaskEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}
