package config

import "os"

func Load() *Config {
	return &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   getEnv("PORT", "8080"),

		// database configuration
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5433"),
		DBUser: getEnv("DB_USER", "kh_user"),
		DBPass: getEnv("DB_PASS", "kh_password"),
		DBName: getEnv("DB_NAME", "knowledge_hub"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
