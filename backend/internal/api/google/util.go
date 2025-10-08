package google

import (
	"log"
	"os"
)

func GetGoogleMapsAPIKey() string {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_MAPS_API_KEY not set")
		return ""
	}
	return apiKey
}
