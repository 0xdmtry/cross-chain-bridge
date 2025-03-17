package routes

import (
	"bridge-accounts-creator/src/config"
	"bridge-accounts-creator/src/controllers/creator_controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, conf *config.Config, creatorController *creator_controller.CreatorController) {
	api := app.Group("v1")

	api.Get("create-account", func(ctx *fiber.Ctx) error {
		return creatorController.CreateAccountController(ctx)
	})
}
