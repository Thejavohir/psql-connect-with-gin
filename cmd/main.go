package main

import (
	"log"
	"net/http"

	"psql/api"
	"psql/config"
	"psql/storage/postgres"
)

func main() {

	cfg := config.Load()

	pgconn, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic("Connection to postgres failed " + err.Error())
	}

	api.NewApi(&cfg, pgconn)

	log.Println("Listening...", cfg.ServerHost+cfg.HttpPort)
	err = http.ListenAndServe(cfg.ServerHost+cfg.HttpPort, nil)
	if err != nil {
		panic("Connection failed to listen to server " + err.Error())
	}

}
