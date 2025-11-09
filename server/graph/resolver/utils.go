package resolver

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
	"workbench/graphql-app/middlewares"
	"workbench/graphql-app/queries/generated"
)

func GenerateID(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.New().String())
}

func MustGetUser(ctx context.Context) (*model.UserClaims, error) {
	claims, ok := ctx.Value(middlewares.AuthCtxKey).(*model.UserClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized: you must be logged in")
	}
	return claims, nil
}

func UniqueStrings(slice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func NewNullString(s *string) sql.NullString {
	if s == nil || *s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func SqlcUserToGraphUser(u generated.User) *model.User {
	var avatar *string
	if u.Avatar.Valid {
		avatar = &u.Avatar.String
	}

	var createdAt time.Time
	if u.CreatedAt.Valid {
		createdAt = u.CreatedAt.Time
	}

	return &model.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Avatar:    avatar,
		CreatedAt: scalar.Date{Time: &createdAt},
	}
}

func SqlcBoardToGraphBoard(b generated.Board) *model.Board {
	return &model.Board{
		ID:   b.ID,
		Name: b.Name,
	}
}

func SqlcColumnToGraphColumn(c generated.Column) *model.Column {
	return &model.Column{
		ID:    c.ID,
		Name:  c.Name,
		Order: int(c.Order),
	}
}

func SqlcTaskToGraphTask(t generated.Task) *model.Task {
	var description *string
	if t.Description.Valid {
		description = &t.Description.String
	}

	var createdAt time.Time
	if t.CreatedAt.Valid {
		createdAt = t.CreatedAt.Time
	}

	return &model.Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: description,
		CreatedAt:   scalar.Date{Time: &createdAt},
		Column:      &model.Column{ID: t.ColumnID},
	}
}

func SqlcCommentToGraphComment(c generated.Comment) *model.Comment {
	var createdAt time.Time
	if c.CreatedAt.Valid {
		createdAt = c.CreatedAt.Time
	}

	return &model.Comment{
		ID:        c.ID,
		Content:   c.Content,
		Author:    &model.User{ID: c.AuthorID},
		Task:      &model.Task{ID: c.TaskID},
		CreatedAt: scalar.Date{Time: &createdAt},
	}
}

func Now() *scalar.Date {
	t := time.Now().UTC()
	return &scalar.Date{Time: &t}
}
