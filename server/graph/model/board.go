package model

type Board struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Columns []*Column `json:"columns"`
}

func (Board) IsNode()         {}
func (b Board) GetID() string { return b.ID }
