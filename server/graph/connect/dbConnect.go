package connect

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"workbench/graphql-app/queries"
	"workbench/graphql-app/queries/generated"
	"workbench/graphql-app/queries/wrapper"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var Queries *wrapper.WrappedQueries

func InitDB() {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if user == "" || pass == "" || name == "" || host == "" || port == "" {
		log.Fatal("Database env vars not set. Set DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_HOST, DB_PORT")
	}

	dsnNoDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true&multiStatements=true", user, pass, host, port)
	tmpDB, err := sql.Open("mysql", dsnNoDB)
	if err != nil {
		log.Fatalf("failed to open mysql (no db): %v", err)
	}
	if err := tmpDB.Ping(); err != nil {
		log.Fatalf("failed to ping mysql (no db): %v", err)
	}

	_, err = tmpDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;", name))
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	_ = tmpDB.Close()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", user, pass, host, port, name)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open mysql: %v", err)
	}

	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	if err := DB.Ping(); err != nil {
		log.Fatalf("failed to ping mysql (db): %v", err)
	}

	// ExecSchema(DB)
	// ExecMock(DB)

	baseQueries := generated.New(DB)
	Queries = queries.InitWrappedQueries(baseQueries)

	log.Println("Database initialized and queries ready")
}

func GetQueries() *wrapper.WrappedQueries {
	if Queries == nil {
		log.Fatal("database not initialized: call connect.InitDB() first")
	}
	return Queries
}
