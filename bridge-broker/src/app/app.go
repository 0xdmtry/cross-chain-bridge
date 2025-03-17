package app

import (
	"bridge-broker/src/config"
	"bridge-broker/src/controllers/broker_controller"
	"bridge-broker/src/routes"
	"bridge-broker/src/services/broker_service"
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

	brokerController := initBroker(config.Conf)

	routes.Setup(app, config.Conf, brokerController)

	return app
}

func initBroker(conf *config.Config) *broker_controller.BrokerController {
	brokerService := broker_service.NewBrokerService(conf)
	brokerController := broker_controller.NewBrokerController(brokerService)
	return brokerController
}
