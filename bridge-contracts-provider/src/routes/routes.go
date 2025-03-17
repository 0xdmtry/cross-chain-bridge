package routes

import (
	"bridge-contracts-provider/src/config"
	"bridge-contracts-provider/src/controllers/provider_controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, conf *config.Config, providerController *provider_controller.ProviderController) {
	api := app.Group("v1")
	api.Get("get-contracts", func(ctx *fiber.Ctx) error {
		return providerController.ProvideContractsController(ctx)
	})
}
