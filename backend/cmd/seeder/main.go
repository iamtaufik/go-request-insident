package main

import (
	"be-request-insident/cmd/seeder/seed"
	"be-request-insident/internal/config"
	"be-request-insident/internal/db"
	"be-request-insident/internal/logger"
	"log"
)

func main() {
	logger.SetupLogger()
	config.LoadEnv()
    mysql := db.ConnectMysql()

	if err := seed.SeedUsers(mysql); err != nil {
		log.Fatal(err)
	}

}