package db

import (
	"log"
	"os"
	"workbench/graphql-app/db/users"

	"github.com/dgraph-io/badger/v4"
)

type Services struct {
	Users *users.Service
}

type Database struct {
	Db       *badger.DB
	Opts     badger.Options
	Services *Services
	Modules  []DbModule

	Stream *badger.Stream
}

const defaultDbPath = "db/data/app.db"

func SetupDB() (*Database, error) {
	dbPath := os.Getenv("BADGER_PATH")
	if dbPath == "" {
		log.Printf("BADGER_PATH not set — using default path: %s", defaultDbPath)
		dbPath = defaultDbPath
	}

	err := os.MkdirAll("db/data", os.ModePerm)
	if err != nil {
		return nil, err
	}

	Database := &Database{
		Opts: badger.DefaultOptions(dbPath),
	}

	if err := Database.initDb(); err != nil {
		return nil, err
	}

	log.Println("BadgerDB initialized successfully")
	return Database, nil
}
