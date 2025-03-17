package transporter_controller

import (
	"bridge-funds-transporter/src/services/transporter_service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type TransporterController struct {
	TransporterService transporter_service.TransporterService
}

func NewTransporterController(service transporter_service.TransporterService) *TransporterController {
	return &TransporterController{
		TransporterService: service,
	}
}

func (c *TransporterController) TransferFundsController(ctx *fiber.Ctx) error {
	fmt.Printf("Transporter::TransporterController::CreateAccountController::ctx: %+v\n", ctx)
	req := new(TransactionPayload)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	if req.ChainUrl == "" || req.PrivateKeyStr == "" || req.RecipientAddress == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
	}

	err := c.TransporterService.TransferFunds(req.ChainUrl, req.PrivateKeyStr, req.RecipientAddress, req.Amount, req.ChainID, req.GasLimit)
	if err != nil {
		return err
	}

	return nil
}
