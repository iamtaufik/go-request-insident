package main

import (
	"log"

	"be-request-insident/internal/config"
	"be-request-insident/internal/db"
	"be-request-insident/internal/handlers"
	"be-request-insident/internal/logger"
	"be-request-insident/internal/repository"
	"be-request-insident/internal/routes"
	"be-request-insident/internal/usecase"

	"github.com/gofiber/fiber/v2"
)



func main() {
    logger.SetupLogger()
	config.LoadEnv()
    app := fiber.New()
    mysql := db.ConnectMysql()
    mongo := db.ConnectMongo()
    redis := db.ConnectRedis()

    appLogger := logger.NewAppLogger(
        mongo.DB.Collection("app_logs"),
        "app_logs",
    )
    userRepo := repository.NewUserRepository(mysql, redis)
    authUsecase := usecase.NewAuthUseCase(userRepo, appLogger)
    authHandler := handlers.NewAuthHandler(authUsecase)

    routes.RegisterRoutes(app, &routes.RouteConfig{
        AuthHandler: authHandler,
    })

    app.Get("/", func (c *fiber.Ctx) error {
        abc := config.GetEnvVariable("ABC")
        log.Println("Value of ABC:", abc)
        return c.SendString("Hello,  world! dsa")
    })

    log.Fatal(app.Listen(":3000"))
}