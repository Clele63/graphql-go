package module

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"workbench/graphql-app/utils"

	"github.com/dgraph-io/badger/v4"
)

func (m *UsersModule) InitValues(bdb *badger.DB) error {
	return bdb.Update(func(txn *badger.Txn) error {
		prefix := []byte(fmt.Sprintf("%s:", m.Namespace()))
		exists := false
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			exists = true
			break
		}

		if exists {
			return nil
		}

		id, err := m.NextID(txn)
		if err != nil {
			return err
		}

		hashedPassword, err := utils.HashPassword("admin")
		if err != nil {
			return err
		}

		admin := map[string]string{
			"id":            strconv.Itoa(id),
			"name":          "admin",
			"password":      hashedPassword,
			"email":         "admin@example.com",
			"creation_date": time.Now().Format("2006-01-02"),
		}

		data, _ := json.Marshal(admin)
		log.Printf("Creating default admin user (id=%d)", id)
		return txn.Set([]byte(fmt.Sprintf("%s:%d", m.Namespace(), id)), data)
	})
}
