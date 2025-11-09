package model

import "workbench/graphql-app/graph/resolver/scalar"

type Comment struct {
	ID        string      `json:"id"`
	Content   string      `json:"content"`
	Author    *User       `json:"author"`
	Task      *Task       `json:"task"`
	CreatedAt scalar.Date `json:"createdAt"`
}

func (Comment) IsNode()         {}
func (c Comment) GetID() string { return c.ID }
