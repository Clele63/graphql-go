package model

type PageInfo struct {
	EndCursor   *string `json:"endCursor,omitempty"`
	HasNextPage bool    `json:"hasNextPage"`
}
