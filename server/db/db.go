package db

import (
	"log"
	"os"

	"go.etcd.io/bbolt"
)

const defaultDbPath = "data/db.go"

type Database struct {
	Db      *bbolt.DB
	Logger  bbolt.DefaultLogger
	DbStats bbolt.Stats
}

func SetupDB() (*Database, error) {
	dbPath := os.Getenv("BBOLT_PATH")
	if dbPath == "" {
		dbPath = defaultDbPath
	}

	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		return nil, err
	}

	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Printf("failed to open BoltDB: %v", err)
		return nil, err
	}

	return &Database{
		Db: db,
	}, nil
}

func (database *Database) InitDB() error {
	err := database.Db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		return err
	})
	return err
}
