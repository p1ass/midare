package config

import (
	"encoding/base64"
	"os"
)

func ReadPort() string {
	return os.Getenv("PORT")
}

func ReadCloudFunctionsURL() string {
	return os.Getenv("CLOUD_FUNCTIONS_URL")
}

func ReadAllowCORSOriginURL() string {
	return os.Getenv("CORS_ALLOW_ORIGIN")
}

func ReadCloudRunRevision() string {
	return os.Getenv("K_REVISION")
}

func ReadFrontEndCallbackURL() string {
	return os.Getenv("FRONTEND_CALLBACK_URL")
}

func ReadSessionKey() string {
	return os.Getenv("SESSION_KEY")
}
func ReadSessionEncryptionKey() ([]byte, error) {
	encoded := os.Getenv("SESSION_ENCRYPTION_KEY_BASE64_ENCODED")
	encryptionKey, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return encryptionKey, nil
}

func ReadGoogleCloudProjectID() string {
	return os.Getenv("GCP_PROJECT")
}

func IsLocal() bool {
	return os.Getenv("ENV") == "LOCAL"
}
