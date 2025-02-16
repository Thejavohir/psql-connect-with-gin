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

	PrivateKey string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     int

	RedisHost     string
	RedisPort     string
	RedisDB       int
	RedisPassword string

	DefaultOffset int
	DefaultLimit  int

	PostgresMaxConnection int32
}

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"

	SuperAdmin = "SUPERADMIN"
)

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Printf(".env file not found %+v", err)
	}

	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", DebugMode))

	c.ServerHost = cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost"))
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost "))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "javohir"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "postgres_connect"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "javohir1"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", ":6379"))
	c.RedisDB = cast.ToInt(getOrReturnDefault("REDIS_DB", 0))
	c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))

	c.PrivateKey = cast.ToString(getOrReturnDefault("PRIVATE_KEY", "samalama1412"))

	c.PostgresMaxConnection = cast.ToInt32(getOrReturnDefault("POSTGRES_MAX_CONNECTION", 30))

	c.DefaultOffset = cast.ToInt(getOrReturnDefault("OFFSET", 0))
	c.DefaultLimit = cast.ToInt(getOrReturnDefault("LIMIT", 10))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
