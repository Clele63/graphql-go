//go:build !embed

package connect

import (
	"log"
	"os"
)

func loadSchema() string {
	schemaPath := "graph/connect/data/schema.sql"
	content, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Printf("warning: cannot read schema file %s: %v", schemaPath, err)
	}
	return string(content)
}

func loadMock() string {
	mockPath := "graph/connect/data/mock.sql"
	content, err := os.ReadFile(mockPath)
	if err != nil {
		log.Printf("warning: cannot read mock file %s: %v", mockPath, err)
	}
	return string(content)
}
