package config

import "os"

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
