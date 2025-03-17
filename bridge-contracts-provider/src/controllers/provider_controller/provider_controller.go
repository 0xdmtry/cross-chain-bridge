package provider_controller

import (
	"bridge-contracts-provider/src/services/provider_service"
	"github.com/gofiber/fiber/v2"
)

type ProviderController struct {
	ProviderService provider_service.ProviderService
}

func NewProviderController(service provider_service.ProviderService) *ProviderController {
	return &ProviderController{
		ProviderService: service,
	}
}

func (c *ProviderController) ProvideContractsController(ctx *fiber.Ctx) error {
	contractsInfo, err := c.ProviderService.ProvideContracts()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return ctx.JSON(contractsInfo)
}
