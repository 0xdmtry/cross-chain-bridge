package account_controller

import (
	"bridge-storage/src/services/account_service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	AccountService account_service.AccountService
}

func NewAccountController(service account_service.AccountService) *AccountController {
	return &AccountController{
		AccountService: service,
	}
}

func (c *AccountController) CreateAccountController(ctx *fiber.Ctx) error {
	fmt.Printf("Storage::AccountController::CreateAccountController::ctx: %+v\n", ctx)
	req := new(CreatAccountRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	if req.PublicKey == "" || req.PrivateKey == "" || req.Address == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
	}

	account, err := c.AccountService.CreateAccountService(req.PublicKey, req.PrivateKey, req.Address)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"account": account,
	})
}
