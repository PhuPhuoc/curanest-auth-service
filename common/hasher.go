package common

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func RandomStr(length int) (string, error) {
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func HashPassword(salt, password string) (string, error) {
	spStr := fmt.Sprintf("%s.%s", salt, password)

	h, err := bcrypt.GenerateFromPassword([]byte(spStr), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

func CompareHashPassword(hashedPassword, salt, password string) bool {
	spStr := fmt.Sprintf("%s.%s", salt, password)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(spStr)) == nil
}
