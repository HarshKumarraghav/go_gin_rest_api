package configs

import (
	"os"
)

var config Config

type Config struct {
	MongoURI  string
	Port      string
	JwtSecret string
}

func FromEnv() Config {
	config = Config{
		MongoURI: os.Getenv("MONGODB_URL"),
		Port:     os.Getenv("PORT"),
	}
	return config
}
