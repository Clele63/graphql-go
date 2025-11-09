package resolver

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
	"workbench/graphql-app/queries/generated"
	"workbench/graphql-app/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	claims, err := MustGetUser(ctx)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	id := GenerateID("u")
	if claims.UserID == id {
		return nil, fmt.Errorf("unauthorized: you cannot modify another user")
	}

	params := generated.CreateUserParams{
		ID:       id,
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := r.Queries.CreateUser(ctx, params); err != nil {
		return nil, fmt.Errorf("unable to create user: %w", err)
	}

	row, err := r.Queries.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve created user: %w", err)
	}

	return SqlcUserToGraphUser(row), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	_, err := MustGetUser(ctx)
	if err != nil {
		return nil, err
	}

	params := generated.UpdateUserParams{
		ID:    input.ID,
		Name:  NewNullString(input.Name),
		Email: NewNullString(input.Email),
	}

	if err := r.Queries.UpdateUser(ctx, params); err != nil {
		return nil, fmt.Errorf("unable to update user: %w", err)
	}

	row, err := r.Queries.GetUser(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve updated user: %w", err)
	}

	return SqlcUserToGraphUser(row), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	claims, err := MustGetUser(ctx)
	if err != nil {
		return false, err
	}

	if claims.UserID == id {
		return false, fmt.Errorf("unauthorized")
	}

	err = r.Queries.DeleteUser(ctx, id)
	return err == nil, err
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	if _, err := MustGetUser(ctx); err != nil {
		return nil, err
	}

	id := GenerateID("t")
	params := generated.CreateTaskParams{
		ID:          id,
		Title:       input.Title,
		Description: NewNullString(input.Description),
		ColumnID:    input.ColumnID,
	}

	if err := r.Queries.CreateTask(ctx, params); err != nil {
		return nil, fmt.Errorf("unable to create task: %w", err)
	}

	if input.AssigneeIds != nil {
		for _, userID := range input.AssigneeIds {
			if userID != "" {
				assignParams := generated.AddTaskAssigneeParams{
					TaskID: id,
					UserID: userID,
				}
				if err := r.Queries.AddTaskAssignee(ctx, assignParams); err != nil {

					return nil, fmt.Errorf("unable to assign user %s: %w", userID, err)
				}
			}
		}
	}

	row, err := r.Queries.GetTask(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve created task: %w", err)
	}

	return SqlcTaskToGraphTask(row), nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	if _, err := MustGetUser(ctx); err != nil {
		return nil, err
	}
	fmt.Println(input)

	params := generated.UpdateTaskParams{
		ID:          input.ID,
		Title:       NewNullString(input.Title),
		Description: NewNullString(input.Description),
	}

	if err := r.Queries.UpdateTask(ctx, params); err != nil {
		return nil, fmt.Errorf("unable to update task: %w", err)
	}

	if input.AssigneeIds != nil {

		if err := r.Queries.ClearTaskAssigneesByTask(ctx, input.ID); err != nil {
			return nil, fmt.Errorf("unable to clear previous assignees: %w", err)
		}

		for _, userID := range input.AssigneeIds {
			if userID != nil {
				assignParams := generated.AddTaskAssigneeParams{
					TaskID: input.ID,
					UserID: *userID,
				}
				if err := r.Queries.AddTaskAssignee(ctx, assignParams); err != nil {
					return nil, fmt.Errorf("unable to assign new user %s: %w", *userID, err)
				}
			}
		}
	}

	row, err := r.Queries.GetTask(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve updated task: %w", err)
	}

	return SqlcTaskToGraphTask(row), nil
}

func (r *mutationResolver) MoveTask(ctx context.Context, id string, toColumnID string) (*model.Task, error) {
	if _, err := MustGetUser(ctx); err != nil {
		return nil, err
	}

	params := generated.MoveTaskParams{
		ID:       id,
		ColumnID: NewNullString(&toColumnID),
	}
	if err := r.Queries.MoveTask(ctx, params); err != nil {
		return nil, fmt.Errorf("unable to move task: %w", err)
	}

	row, err := r.Queries.GetTask(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve moved task: %w", err)
	}

	return SqlcTaskToGraphTask(row), nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	if _, err := MustGetUser(ctx); err != nil {
		return false, err
	}

	err := r.Queries.DeleteTask(ctx, id)
	return err == nil, err
}

func (r *mutationResolver) AddComment(ctx context.Context, taskID string, content string) (*model.Comment, error) {
	claims, err := MustGetUser(ctx)
	if err != nil {
		return nil, err
	}

	id := GenerateID("cm")
	params := generated.CreateCommentParams{
		ID:       id,
		Content:  content,
		AuthorID: claims.UserID,
		TaskID:   taskID,
	}

	if err := r.Queries.CreateComment(ctx, params); err != nil {
		return nil, fmt.Errorf("unable to add comment: %w", err)
	}

	row, err := r.Queries.GetComment(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve created comment: %w", err)
	}

	return SqlcCommentToGraphComment(row), nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	claims, err := MustGetUser(ctx)
	if err != nil {
		return false, err
	}

	authorID, err := r.Queries.GetCommentAuthor(ctx, id)
	if err != nil {
		return false, fmt.Errorf("comment not found: %w", err)
	}

	if claims.UserID != authorID {

		return false, fmt.Errorf("unauthorized: you are not the author")
	}

	err = r.Queries.DeleteComment(ctx, id)
	return err == nil, err
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.Token, error) {
	row, err := r.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("incorrect email or password")
		}

		return nil, fmt.Errorf("database error: %w", err)
	}

	if !utils.ComparePassword(password, row.Password) {
		return nil, fmt.Errorf("incorrect email or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	tokenString := utils.GenerateJwt(row.ID, expirationTime.Unix())

	return &model.Token{
		Token:     tokenString,
		ExpiredAt: scalar.Date{Time: &expirationTime},
	}, nil
}
