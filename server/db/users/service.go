package users

import (
	"strconv"
	"time"
	"workbench/graphql-app/utils"

	"github.com/dgraph-io/badger/v4"
)

type Service struct {
	store *Store
}

func NewService(db *badger.DB) *Service {
	return &Service{
		store: NewStore(db),
	}
}

// CreateUser crée un nouvel utilisateur avec ID auto-incrémenté et mot de passe hashé
func (s *Service) CreateUser(name, password, email string) (map[string]string, error) {
	var newUser map[string]string

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	err = s.store.db.Update(func(txn *badger.Txn) error {
		id, err := s.store.module.NextID(txn)
		if err != nil {
			return err
		}

		user := map[string]string{
			"id":            strconv.Itoa(id),
			"name":          name,
			"password":      hashedPassword,
			"email":         email,
			"creation_date": time.Now().Format("2006-01-02"),
		}

		newUser = user
		return s.store.PutUser(user)
	})
	return newUser, err
}

// UpdateUser met à jour un utilisateur (champs optionnels)
func (s *Service) UpdateUser(id string, name *string) (map[string]string, error) {
	var updatedUser map[string]string

	// On doit utiliser Update car GetUserByID est en View-only
	err := s.store.db.Update(func(txn *badger.Txn) error {
		// 1. Récupérer l'utilisateur existant
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		user, err := s.store.GetUserByID(idInt) // store.GetUserByID utilise un int
		if err != nil {
			return err
		}

		// 2. Mettre à jour les champs si fournis
		if name != nil {
			user["name"] = *name
		}

		// 3. Sauvegarder l'utilisateur
		updatedUser = user
		return s.store.PutUser(user)
	})
	return updatedUser, err
}

func (s *Service) DeleteUser(id string) error {
	return s.store.DeleteUser(id)
}

func (s *Service) GetUserByID(id int) (map[string]string, error) {
	return s.store.GetUserByID(id)
}

func (s *Service) GetUserAuthByName(name string) (map[string]string, error) {
	return s.store.GetUserByName(name)
}

func (s *Service) ListUsers() ([]map[string]string, error) {
	return s.store.ListUsers()
}

func (s *Service) SearchUsersByName(term string) ([]map[string]string, error) {
	return s.store.SearchUsersByName(term)
}
