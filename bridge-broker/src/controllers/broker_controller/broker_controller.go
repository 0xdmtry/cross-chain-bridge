package broker_controller

import (
	"bridge-broker/src/services/broker_service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type BrokerController struct {
	BrokerService broker_service.BrokerService
}

func NewBrokerController(service broker_service.BrokerService) *BrokerController {
	return &BrokerController{
		BrokerService: service,
	}
}

func (c *BrokerController) CreateAccountController(ctx *fiber.Ctx) error {
	fmt.Printf("\nfmt::Broker::CreateAccountController\n")
	account, err := c.BrokerService.CreateAccountService()

	if err != nil {
		fmt.Printf("\nfmt::Broker::CreateAccountController::ERROR: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	fmt.Printf("\nfmt::Broker::CreateAccountController::CreatedAccount: %v\n", account)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})
}
