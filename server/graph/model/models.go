package model

type Node interface {
	IsNode()
	GetID() string
}

type SearchResult interface {
	IsSearchResult()
}

type CreateTaskInput struct {
	Title       string   `json:"title"`
	Description *string  `json:"description,omitempty"`
	AssigneeIds []string `json:"assigneeIds,omitempty"`
	ColumnID    string   `json:"columnId"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Mutation struct {
}

type Query struct {
}

type Subscription struct {
}

type UpdateTaskInput struct {
	ID          string    `json:"id"`
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	AssigneeIds []*string `json:"assigneeIds,omitempty"`
}

type UpdateUserInput struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
