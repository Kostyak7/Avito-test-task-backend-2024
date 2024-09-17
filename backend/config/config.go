package config

import (
    "os"
)

type Config struct {
    ServerAddress    string
    PostgresConn     string
    PostgresJDBCURL  string
    PostgresUsername string
    PostgresPassword string
    PostgresHost     string
    PostgresPort     string
    PostgresDatabase string
}

func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}

func LoadConfig() *Config {
    return &Config{
        ServerAddress:    getEnv("SERVER_ADDRESS", "0.0.0.0:8080"),
        PostgresConn:     getEnv("POSTGRES_CONN", ""),
        PostgresJDBCURL:  getEnv("POSTGRES_JDBC_URL", ""),
        PostgresUsername: getEnv("POSTGRES_USERNAME", ""),
        PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),
        PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
        PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
        PostgresDatabase: getEnv("POSTGRES_DATABASE", ""),
    }
}
