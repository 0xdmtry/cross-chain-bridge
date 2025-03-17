package compiler_controller

import (
	"bridge-eth-compiler/src/helpers/logger"
	"bridge-eth-compiler/src/services/compiler_service"
	"github.com/gofiber/fiber/v2"
)

type CompilerController struct {
	CompilerService compiler_service.CompilerService
}

func NewCompilerController(service compiler_service.CompilerService) *CompilerController {
	return &CompilerController{
		CompilerService: service,
	}
}

func (c *CompilerController) CompileContractController(ctx *fiber.Ctx) error {
	path := ctx.Query("path")
	output := ctx.Query("output")
	name := ctx.Query("name")

	outputPath, err := c.CompilerService.CompileContract(path, output, name)
	if err != nil {
		logger.Error("EthCompiler::CompileContractController:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.JSON(outputPath)
}
