package main

import (
	"log"

	"be-request-insident/internal/config"
	"be-request-insident/internal/db"
	"be-request-insident/internal/logger"

	"github.com/gofiber/fiber/v2"
)



func main() {
    logger.SetupLogger()
	config.LoadEnv()
    app := fiber.New()
    db.ConnectMysql()
    db.ConnectMongo()
    db.ConnectRedis()

    app.Get("/", func (c *fiber.Ctx) error {
        abc := config.GetEnvVariable("ABC")
        log.Println("Value of ABC:", abc)
        return c.SendString("Hello,  world! dsa")
    })

    log.Fatal(app.Listen(":3000"))
}