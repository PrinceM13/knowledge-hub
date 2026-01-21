package config

type Config struct {
	AppEnv string
	Port   string

	// database configuration
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}
