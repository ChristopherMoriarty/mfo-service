package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	DB     DB
	HTTP   HTTPConfig
	Logger LoggerConfig
}

type DB struct {
	PG_Host     string `env:"DB_HOST"     envDefault:"localhost"`
	PG_Port     int    `env:"DB_PORT"     envDefault:"5432"`
	PG_User     string `env:"DB_USER"     envRequired:"true"`
	PG_Password string `env:"DB_PASS"     envRequired:"true"`
	PG_Name     string `env:"DB_NAME"     envRequired:"true"`
	PG_SSLMode  string `env:"DB_SSLMODE"  envDefault:"disable"`
}

func (db *DB) GetURI() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		db.PG_User,
		db.PG_Password,
		db.PG_Host,
		db.PG_Port,
		db.PG_Name,
		db.PG_SSLMode,
	)
}

type HTTPConfig struct {
	Port int `env:"PORT" envDefault:"8080"`
}

type LoggerConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"info"`
}

func (c *LoggerConfig) GetLevel() zapcore.Level {
	switch c.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Parse() (*Config, error) {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Printf(".env file not found or could not be loaded: %v\n", err)
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("load config from env: %w", err)
	}
	return &cfg, nil
}
