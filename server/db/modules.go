package db

import "github.com/dgraph-io/badger/v4"

// DBModule représente un module indépendant qui peut initialiser une partie de la base
type DbModule interface {
	Namespace() string
	InitSchema(db *badger.DB) error
	InitValues(db *badger.DB) error
	NextID(txn *badger.Txn) (int, error)
}

// RegisterModule ajoute un module au registre global
func (database *Database) RegisterModule(m DbModule) {
	database.Modules = append(database.Modules, m)
}

// initAllModules exécute l’initialisation de tous les modules enregistrés
func (database *Database) initAllModules() error {
	for _, module := range database.Modules {
		if err := module.InitSchema(database.Db); err != nil {
			return err
		}
		if err := module.InitValues(database.Db); err != nil {
			return err
		}
	}
	return nil
}
