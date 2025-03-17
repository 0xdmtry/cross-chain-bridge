package deployer_controller

import (
	"bridge-eth-deployer/src/helpers/logger"
	"bridge-eth-deployer/src/services/deployer_service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type DeployerController struct {
	DeployerService deployer_service.DeployerService
}

func NewDeployerController(service deployer_service.DeployerService) *DeployerController {
	return &DeployerController{
		DeployerService: service,
	}
}

func (c *DeployerController) DeployContractController(ctx *fiber.Ctx) error {
	path := ctx.FormValue("path")
	name := ctx.FormValue("name")
	endpoint := ctx.FormValue("endpoint")
	walletKey := ctx.FormValue("walletKey")
	chainId := ctx.FormValue("chainId")

	id, err := strconv.ParseInt(chainId, 10, 64)
	if err != nil {
		logger.Error("EthDeployer::DeployContractController::strconv.ParseInt:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = c.DeployerService.DeployContract(path, name, endpoint, walletKey, id)
	if err != nil {
		logger.Error("EthDeployer::DeployContractController::DeployContract:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return nil
}
