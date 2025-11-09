package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash original password
func HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pass), err
}

// HashPasswordList hashes each password in the list
func HashPasswordList(passwords []string) ([]string, error) {
	hashes := make([]string, len(passwords))

	for i, p := range passwords {
		h, err := HashPassword(p)
		if err != nil {
			return nil, err
		}
		hashes[i] = h
	}

	return hashes, nil
}

// ComparePassword returns true if password matches hash
func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
