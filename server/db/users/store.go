package users

import (
	"encoding/json"
	"fmt"
	"strings"
	"workbench/graphql-app/db/users/module"

	"github.com/dgraph-io/badger/v4"
)

type Store struct {
	db     *badger.DB
	module module.UsersModule
}

func NewStore(db *badger.DB) *Store {
	return &Store{db: db}
}

// GetUserByID lit un utilisateur depuis Badger
func (s *Store) GetUserByID(id int) (map[string]string, error) {
	var user map[string]string

	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("%s:%d", s.module.Namespace(), id)))
		if err != nil {
			return err
		}
		return item.Value(func(v []byte) error {
			return json.Unmarshal(v, &user)
		})
	})
	return user, err
}

// GetUserByName recherche un utilisateur par son nom (pour le login)
func (s *Store) GetUserByName(name string) (map[string]string, error) {
	var user map[string]string
	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(fmt.Sprintf("%s:", s.module.Namespace()))
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			var currentUser map[string]string
			err := item.Value(func(v []byte) error {
				return json.Unmarshal(v, &currentUser)
			})
			if err != nil {
				return err // ou continuer ? Pour l'instant, on arrête.
			}

			if currentUser["name"] == name {
				user = currentUser
				return nil // Utilisateur trouvé
			}
		}
		return badger.ErrKeyNotFound // Utilisateur non trouvé après itération
	})

	return user, err
}

// PutUser enregistre un utilisateur (création ou update)
func (s *Store) PutUser(u map[string]string) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		key := []byte(fmt.Sprintf("%s:%s", s.module.Namespace(), u["id"]))
		return txn.Set(key, data)
	})
}

// DeleteUser supprime un utilisateur
func (s *Store) DeleteUser(id string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		key := []byte(fmt.Sprintf("%s:%s", s.module.Namespace(), id))
		return txn.Delete(key)
	})
}

// ListUsers renvoie tous les utilisateurs
func (s *Store) ListUsers() ([]map[string]string, error) {
	users := []map[string]string{}

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(fmt.Sprintf("%s:", s.module.Namespace()))
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			var user map[string]string
			err := item.Value(func(v []byte) error {
				return json.Unmarshal(v, &user)
			})
			if err != nil {
				return err
			}
			users = append(users, user)
		}
		return nil
	})
	return users, err
}

// SearchUsersByName renvoie les utilisateurs dont le nom contient le terme de recherche (insensible à la casse)
func (s *Store) SearchUsersByName(term string) ([]map[string]string, error) {
	users := []map[string]string{}
	lowerTerm := strings.ToLower(term)

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		prefix := []byte(fmt.Sprintf("%s:", s.module.Namespace()))
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			var user map[string]string
			err := item.Value(func(v []byte) error {
				return json.Unmarshal(v, &user)
			})
			if err != nil {
				return err
			}

			if strings.Contains(strings.ToLower(user["name"]), lowerTerm) {
				users = append(users, user)
			}
		}
		return nil
	})
	return users, err
}
