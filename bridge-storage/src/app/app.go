package app

import (
	"bridge-storage/src/config"
	"bridge-storage/src/controllers/account_controller"
	"bridge-storage/src/databases/mysql"
	account_dao "bridge-storage/src/models/account_model/dao"
	"bridge-storage/src/routes"
	"bridge-storage/src/services/account_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func New() *fiber.App {
	app := fiber.New()
	db := mysql.GetDB()

	accountController := initAccount(db, config.Conf)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, UPDATE, DELETE, OPTIONS",
		AllowHeaders: "*",
	}))

	routes.Setup(app, config.Conf, accountController)

	return app
}

func initAccount(db *gorm.DB, conf *config.Config) *account_controller.AccountController {
	accountDAO := account_dao.NewAccountDAO(db)
	accountService := account_service.NewAccountService(accountDAO, conf)
	accountController := account_controller.NewAccountController(accountService)
	return accountController
}
