package db

import (
	"be-request-insident/internal/config"
	"context"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	redisAddr := config.GetEnvVariable("REDIS_ADDR")
	redisPort := config.GetEnvVariable("REDIS_PORT")
	redisDb := config.GetEnvVariable("REDIS_DB")
	db, err := strconv.Atoi(redisDb)

	if err != nil {
		db = 0 // default to DB 0 if conversion fails
	}

	rdb := redis.NewClient(&redis.Options{
        Addr:     redisAddr + ":" + redisPort,
        Password: "", // no password set
        DB:       db,  
    })

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to Redis successfully")
	return rdb
}