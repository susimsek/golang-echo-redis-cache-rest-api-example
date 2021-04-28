package config

import "os"

var (
	ServerPort            = GetEnv("SERVER_PORT", "9000")
	RedisUrl              = GetEnv("REDIS_URL", "localhost:6379")
	RedisPassword         = GetEnv("REDIS_PASSWORD", "")
	UserCacheExpirationMs = GetEnv("USER_CACHE_EXPIRATION_MS", "3600000")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
