package app

import (
	"bridge-funds-transporter/src/config"
	"bridge-funds-transporter/src/controllers/transporter_controller"
	"bridge-funds-transporter/src/routes"
	"bridge-funds-transporter/src/services/transporter_service"
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

	transporterController := initTransporter(config.Conf)

	routes.Setup(app, config.Conf, transporterController)

	return app
}

func initTransporter(conf *config.Config) *transporter_controller.TransporterController {
	transporterService := transporter_service.NewTransporterService(conf)
	transporterController := transporter_controller.NewTransporterController(transporterService)
	return transporterController
}
