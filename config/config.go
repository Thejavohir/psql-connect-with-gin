package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	ServerHost string
	HttpPort   string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     int

	DefaultOffset int
	DefaultLimit int

	PostgresMaxConnection int32
}

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"
)

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Printf(".env file not found %+v", err)
	}

	cfg := Config{}

	cfg.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST", "localhost"))
	cfg.HttpPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost "))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "javohir"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "postgres_connect"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "javohir1"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))

	cfg.PostgresMaxConnection = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTION", 30))
	
	cfg.DefaultOffset = cast.ToInt(getOrReturnDefaultValue("OFFSET", 0))
	cfg.DefaultLimit = cast.ToInt(getOrReturnDefaultValue("LIMIT", 10))

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
