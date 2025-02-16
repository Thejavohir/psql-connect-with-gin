package main

import (
	"psql/api"
	"psql/config"
	"psql/pkg/logger"
	"psql/storage/postgres"
	"psql/storage/redis"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	var loggerLevel = new(string)

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	logr := logger.New(*loggerLevel, "app")
	defer func() {
		err := logger.Cleanup(logr)
		if err != nil {
			return
		}
	}()

	pgconn, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic("Connection to postgres failed " + err.Error())
	}
	defer pgconn.Close()

	redisconn, err := redis.NewConnectionRedis(&cfg)
	if err != nil {
		panic("Connectin to redis failed " + err.Error())
	}
	defer redisconn.Close()

	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	api.NewApi(r, &cfg, pgconn, logr, redisconn)

	err = r.Run(cfg.ServerHost + cfg.HttpPort)
	if err != nil {
		panic("Connection failed to listen to server " + err.Error())
	}

}
