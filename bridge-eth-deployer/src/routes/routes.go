package routes

import (
	"bridge-eth-deployer/src/config"
	"bridge-eth-deployer/src/controllers/deployer_controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, conf *config.Config, deployerController *deployer_controller.DeployerController) {
	api := app.Group("v1")

	api.Post("deploy-contract", func(ctx *fiber.Ctx) error {
		return deployerController.DeployContractController(ctx)
	})
}
