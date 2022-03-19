package config

import "os"

func ReadDatastoreProjectId() string {
	return os.Getenv("DATASTORE_PROJECT_ID")
}
