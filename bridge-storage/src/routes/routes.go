package routes

import (
	"bridge-storage/src/config"
	"bridge-storage/src/controllers/account_controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, conf *config.Config, accountController *account_controller.AccountController) {
	api := app.Group("v1")

	api.Post("create-account", func(ctx *fiber.Ctx) error {
		return accountController.CreateAccountController(ctx)
	})
}
