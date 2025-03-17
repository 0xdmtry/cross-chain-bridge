package creator_controller

import (
	"bridge-accounts-creator/src/services/creator_service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type CreatorController struct {
	CreatorService creator_service.CreatorService
}

func NewCreatorController(service creator_service.CreatorService) *CreatorController {
	return &CreatorController{
		CreatorService: service,
	}
}

func (c *CreatorController) CreateAccountController(ctx *fiber.Ctx) error {
	fmt.Printf("\nfmt::Creator::CreateAccountController\n")
	account, err := c.CreatorService.CreateAccountService()
	if err != nil {
		fmt.Printf("\nfmt::Creator::CreateAccountController::ERROR: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	fmt.Printf("\nfmt::Creator::CreateAccountController::Account: %v\n", account)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})
}
