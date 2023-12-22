package utility

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	saltSize := 10
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password+string(salt)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(append(salt, hash...)), nil
}

func VerifyPassword(hashedPassword string, password string) (bool, error) {
	saltSize := 10
	hash := hashedPassword

	if len(hash) < bcrypt.DefaultCost {
		return false, errors.New("invalid hash")
	}

	// Extract the salt from the hash
	if len(hash) < saltSize+bcrypt.MinCost {
		return false, errors.New("invalid hash")
	}
	salt := hash[:saltSize]
	hashWithoutSalt := hash[saltSize:]

	if err := bcrypt.CompareHashAndPassword([]byte(hashWithoutSalt), []byte(password+string(salt))); err != nil {
		return false, err
	}

	return true, nil
}
