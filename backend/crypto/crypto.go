package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// LongSecureRandomBase64 returns 64byte length secure random string
func LongSecureRandomBase64() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// ShortSecureRandomBase64 returns 32byte length secure random string
func ShortSecureRandomBase64() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
