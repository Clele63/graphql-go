package connect

import (
	"database/sql"
	"log"
)

func ExecSchema(DB *sql.DB) {
	schemaSQL := loadSchema()
	if _, err := DB.Exec(schemaSQL); err != nil {
		log.Fatalf("failed to execute schema.sql: %v", err)
	}
}

func ExecMock(DB *sql.DB) {
	mockSQL := loadMock()
	if _, err := DB.Exec(mockSQL); err != nil {
		log.Fatalf("failed to execute mock.sql: %v", err)
	}
}
