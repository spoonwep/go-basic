package utils

import "os"

func IsLocal() bool {
	if os.Getenv("APP_ENV") == "local" {
		return true
	}
	return false
}

func IsDevelopment() bool {
	if os.Getenv("APP_ENV") == "development" {
		return true
	}
	return false
}

func IsProduction() bool {
	if os.Getenv("APP_ENV") == "production" {
		return true
	}
	return false
}
