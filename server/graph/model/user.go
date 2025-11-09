package model

import (
	"workbench/graphql-app/graph/resolver/scalar"
)

type User struct {
	ID        string      `json:"id" gorm:"primary_key"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Avatar    *string     `json:"avatar,omitempty"`
	CreatedAt scalar.Date `json:"created_at"`
}

func (User) IsNode()         {}
func (u User) GetID() string { return u.ID }
func (User) IsSearchResult() {}
