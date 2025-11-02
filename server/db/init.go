package db

import (
	"log"
	"time"
	"workbench/graphql-app/db/users/module"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
)

func (database *Database) initDbOpts() {
	opts := &database.Opts
	opts.SyncWrites = true
	opts.NumVersionsToKeep = 1
	opts.CompactL0OnClose = true
	opts.Compression = options.ZSTD
	opts.ZSTDCompressionLevel = 3
	opts.NumMemtables = 5
	opts.NumLevelZeroTables = 2
	opts.NumLevelZeroTablesStall = 10
	opts.ReadOnly = false
	opts.ValueLogFileSize = 1 << 30
}

func (database *Database) initDbMeta() error {
	return database.Db.Update(func(txn *badger.Txn) error {
		// Exemple : stocker une clé d’initialisation si elle n’existe pas
		_, err := txn.Get([]byte("meta:init"))
		if err == badger.ErrKeyNotFound {
			log.Println("First-time DB setup: creating meta:init flag and default metadata")

			if err := txn.Set([]byte("meta:init"), []byte(time.Now().Format(time.RFC3339))); err != nil {
				return err
			}

			// Exemple de valeur par défaut
			if err := txn.Set([]byte("app:version"), []byte("1.0.0")); err != nil {
				return err
			}
		}

		return nil
	})
}

func (database *Database) openDb() error {
	var err error
	database.Db, err = badger.Open(database.Opts)
	if err != nil {
		log.Fatalf("failed to open BadgerDB: %v", err)
		return err
	}
	return database.initDbMeta()
}

func (database *Database) initDb() error {
	database.initDbOpts()

	if err := database.openDb(); err != nil {
		return err
	}

	database.RegisterModule(&module.UsersModule{})

	if err := database.initAllModules(); err != nil {
		return err
	}
	return nil
}
