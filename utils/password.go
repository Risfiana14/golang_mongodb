package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword -> buat hash dari password
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", bcrypt.ErrMismatchedHashAndPassword
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash -> verifikasi password dengan hash dari DB
func CheckPasswordHash(password, hash string) bool {
	if password == "" || hash == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
