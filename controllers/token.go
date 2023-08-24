package controllers

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/google/uuid"
)

// randToken generates a random hex value.
func GetToken() (string, error) {
	n := 20 // total character
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateUUID() string {
	return uuid.NewString()
}
