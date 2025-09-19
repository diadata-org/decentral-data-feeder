package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// GetImageVersion returns the Docker image version from environment variable
func GetImageVersion() string {
	// Get version from IMAGE_TAG environment variable
	version := os.Getenv("IMAGE_TAG")

	if version == "" {
		version = "unknown" // fallback if not set
		log.Info("No version found, using 'unknown'")
	}

	return version
}
