package app

import (
	"bridge-contracts-provider/src/config"
	"bridge-contracts-provider/src/controllers/provider_controller"
	"bridge-contracts-provider/src/models/dao"
	"bridge-contracts-provider/src/routes"
	"bridge-contracts-provider/src/services/provider_service"
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

	providerController := initProvider(config.Conf)

	routes.Setup(app, config.Conf, providerController)

	return app
}

func initProvider(conf *config.Config) *provider_controller.ProviderController {
	contractModel := dao.NewContractDAO(conf)
	providerService := provider_service.NewProviderService(conf, contractModel)
	providerController := provider_controller.NewProviderController(providerService)
	return providerController
}
