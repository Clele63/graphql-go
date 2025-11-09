package model

type Column struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Order int             `json:"order"`
	Tasks *TaskConnection `json:"tasks"`
}

func (Column) IsNode()         {}
func (c Column) GetID() string { return c.ID }
