package app

import (
	"bridge-accounts-creator/src/config"
	"bridge-accounts-creator/src/controllers/creator_controller"
	"bridge-accounts-creator/src/routes"
	"bridge-accounts-creator/src/services/creator_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, UPDATE, DELETE, OPTIONS",
		AllowHeaders: "*",
	}))

	deployerController := initEOACreator(config.Conf)

	routes.Setup(app, config.Conf, deployerController)

	return app
}

func initEOACreator(conf *config.Config) *creator_controller.CreatorController {
	creatorService := creator_service.NewCreatorService(conf)
	creatorController := creator_controller.NewCreatorController(creatorService)
	return creatorController
}
