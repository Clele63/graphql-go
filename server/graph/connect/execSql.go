package connect

import (
	"database/sql"
	"log"
	"strings"
	"workbench/graphql-app/utils"
)

func ExecSchema(DB *sql.DB) {
	schemaSQL := loadSchema()
	if _, err := DB.Exec(schemaSQL); err != nil {
		log.Fatalf("failed to execute schema.sql: %v", err)
	}
}

func ExecMock(DB *sql.DB) {
	mockSQL, err := replaceIntoSql(loadMock())
	if err != nil {
		return
	}

	if _, err := DB.Exec(mockSQL); err != nil {
		log.Fatalf("failed to execute mock.sql: %v", err)
	}
}

func replaceIntoSql(sqlLoaded string) (string, error) {
	passwords := []string{
		"password",
	}

	hashes, err := utils.HashPasswordList(passwords)
	if err != nil {
		log.Panicf("failed to hash passwords: %v", err)
		return "", err
	}

	replacements := map[string]string{
		"{{pwd_password}}": hashes[0],
	}

	for placeholder, val := range replacements {
		sqlLoaded = strings.ReplaceAll(sqlLoaded, placeholder, val)
	}

	return sqlLoaded, nil
}
