package routes

import (
	"bridge-broker/src/config"
	"bridge-broker/src/controllers/broker_controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, conf *config.Config, brokerController *broker_controller.BrokerController) {
	api := app.Group("v1")
	api.Get("create-account", func(ctx *fiber.Ctx) error {
		return brokerController.CreateAccountController(ctx)
	})
}
