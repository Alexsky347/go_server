package config

import (
	env "github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"path/filepath"
	"runtime"
	"time"
)

const (
	ProductionEnv = "production"

	DatabaseTimeout    = 5 * time.Second
	ProductCachingTime = 1 * time.Minute
)

var AuthIgnoreMethods = []string{
	"/user.UserService/Login",
	"/user.UserService/Register",
}

type Schema struct {
	Environment      string `env:"environment"`
	LogLevel         string `env:"log_level"`
	HTTPPort         int    `env:"http_port"`
	Hostname         string `env:"hostname"`
	AuthSecret       string `env:"auth_secret"`
	Dsn              string `env:"database_url"`
	RedisURI         string `env:"redis_uri"`
	RedisPassword    string `env:"redis_password"`
	RedisDB          int    `env:"redis_db"`
	DefaultStaticDir string `env:"default_static_dir"`
}

var (
	cfg Schema
)

func LoadConfig() (*Schema, error) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	err := godotenv.Load(filepath.Join(currentDir, "config.yml"))
	if err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Schema {
	return &cfg
}
