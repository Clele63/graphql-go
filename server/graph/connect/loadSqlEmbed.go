//go:build embed

package connect

import (
	_ "embed"
)

//go:embed data/schema.sql
var schemaSQL string

// //go:embed data/mock.sql
var mockSQL string

func loadSchema() string {
	return schemaSQL
}

func loadMock() string {
	return mockSQL
}
