package routes

import (
	"bridge-funds-transporter/src/config"
	"bridge-funds-transporter/src/controllers/transporter_controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, conf *config.Config, transporterController *transporter_controller.TransporterController) {
	api := app.Group("v1")

	api.Post("transfer-funds", func(ctx *fiber.Ctx) error {
		return transporterController.TransferFundsController(ctx)
	})

}
