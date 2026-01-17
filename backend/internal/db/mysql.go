package db

import (
	"be-request-insident/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectMysql() *gorm.DB {
	host := config.GetEnvVariable("MYSQL_HOST")
	port := config.GetEnvVariable("MYSQL_PORT")
	dbName := config.GetEnvVariable("MYSQL_DATABASE")
	dbUser := config.GetEnvVariable("MYSQL_USER")
	dbPassword := config.GetEnvVariable("MYSQL_PASSWORD")

	MYSQL_DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		dbUser, dbPassword, host, port, dbName)

	db, err := gorm.Open(mysql.Open(MYSQL_DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	log.Println("Connected to MySQL successfully")
	return db
}