package app

import (
	"bridge-eth-deployer/src/config"
	"bridge-eth-deployer/src/controllers/deployer_controller"
	"bridge-eth-deployer/src/routes"
	"bridge-eth-deployer/src/services/deployer_service"
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

	deployerController := initDeployer(config.Conf)

	routes.Setup(app, config.Conf, deployerController)

	return app
}

func initDeployer(conf *config.Config) *deployer_controller.DeployerController {
	deployerService := deployer_service.NewDeployerService(conf)
	deployerController := deployer_controller.NewDeployerController(deployerService)
	return deployerController
}
