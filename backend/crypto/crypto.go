package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// SecureRandomBase64Encoded returns base 64 url encoded secure random string
func SecureRandomBase64Encoded(entropyByte int) string {
	b := make([]byte, entropyByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
