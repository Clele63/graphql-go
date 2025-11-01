package model

import "workbench/graphql-app/graph/resolver/scalar"

type Node interface {
	IsNode()
	GetID() string
}

type SearchResult interface {
	IsSearchResult()
}

type CreateUserInput struct {
	Name         string      `json:"name"`
	Password     string      `json:"password"`
	Email        string      `json:"email"`
	CreationDate scalar.Date `json:"creation_date"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateUserInput struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
