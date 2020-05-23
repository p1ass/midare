package crypto

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// SecureRandom returns a random uuid string
func SecureRandom() string {
	return uuid.New().String()
}

// SecureRandomBase64 returns a random uuid string encoded by base64
func SecureRandomBase64() string {
	return base64.StdEncoding.EncodeToString(uuid.New().NodeID())
}

// LongSecureRandomBase64 returns twice length secure random string
func LongSecureRandomBase64() string {
	return SecureRandomBase64() + SecureRandomBase64()
}
