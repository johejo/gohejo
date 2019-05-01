package envutils

import "os"

// GetEnv wraps os.Getenv with default value.
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) < 1 {
		value = defaultValue
	}
	return value
}
