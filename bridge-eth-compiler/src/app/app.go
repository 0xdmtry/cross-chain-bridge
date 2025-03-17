package app

import (
	"bridge-eth-compiler/src/config"
	"bridge-eth-compiler/src/controllers/compiler_controller"
	"bridge-eth-compiler/src/routes"
	"bridge-eth-compiler/src/services/compiler_service"
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

	compilerController := initCompiler(config.Conf)

	routes.Setup(app, config.Conf, compilerController)

	return app
}

func initCompiler(conf *config.Config) *compiler_controller.CompilerController {
	compilerService := compiler_service.NewCompilerService(conf)
	compilerController := compiler_controller.NewCompilerController(compilerService)
	return compilerController
}
