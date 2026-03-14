package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv      string
	AppPort     string
	DBType      string
	DBPath      string
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPass      string
	JWTSecret   string
	UploadPath  string
	MaxUploadMB int
}

func Load() *Config {
	for _, path := range []string{".env", "../.env"} {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}
	return &Config{
		AppEnv:      getEnv("APP_ENV", "dev"),
		AppPort:     getEnv("APP_PORT", "8080"),
		DBType:      getEnv("DB_TYPE", "sqlite"),
		DBPath:      getEnv("DB_PATH", "./data/dev.db"),
		DBHost:      getEnv("DB_HOST", ""),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBName:      getEnv("DB_NAME", "metal_db"),
		DBUser:      getEnv("DB_USER", ""),
		DBPass:      getEnv("DB_PASS", ""),
		JWTSecret:   getEnv("JWT_SECRET", "change_me_in_prod"),
		UploadPath:  getEnv("UPLOAD_PATH", "./storage"),
		MaxUploadMB: getEnvInt("MAX_UPLOAD_SIZE_MB", 10),
	}
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultVal
	}
	return i
}
