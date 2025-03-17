package routes

import (
	"bridge-eth-compiler/src/config"
	"bridge-eth-compiler/src/controllers/compiler_controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, conf *config.Config, compilerController *compiler_controller.CompilerController) {
	api := app.Group("v1")
	api.Get("compile-contract", func(ctx *fiber.Ctx) error {
		return compilerController.CompileContractController(ctx)
	})
}
