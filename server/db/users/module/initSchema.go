package module

import (
	"log"

	"github.com/dgraph-io/badger/v4"
)

func (m *UsersModule) InitSchema(bdb *badger.DB) error {
	return bdb.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte(m.NamespaceSchema())); err == badger.ErrKeyNotFound {
			log.Printf("Initializing %s schema", m.Namespace())
			if err := txn.Set([]byte(m.NamespaceSchema()), []byte("initialized")); err != nil {
				return err
			}
			if err := txn.Set([]byte(m.NamespaceSeq()), []byte("0")); err != nil {
				return err
			}
		}
		return nil
	})
}
