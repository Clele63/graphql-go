package connect

import (
	"log"

	"workbench/graphql-app/db"
	"workbench/graphql-app/db/users"
)

// InitDB opens the connection, ensures the database exists, runs the schema and inits sqlc Queries.
func InitDB() *db.Database {
	BadgerDB, err := db.SetupDB()
	if err != nil {
		log.Fatalf("failed to setup the Badger Db : %v", err)
	}

	BadgerDB.Services = &db.Services{
		Users: users.NewService(BadgerDB.Db),
	}

	log.Println("Database and Services initialized. Queries of services are ready")
	return BadgerDB
}
