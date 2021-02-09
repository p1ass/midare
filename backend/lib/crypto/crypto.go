package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// LongSecureRandomBase64 returns twice length secure random string
func LongSecureRandomBase64() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
