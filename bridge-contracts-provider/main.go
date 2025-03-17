package main

import (
	"bridge-contracts-provider/src/app"
	"bridge-contracts-provider/src/config"
	"bridge-contracts-provider/src/helpers/logger"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

func main() {
	config.New()
	err := logger.Initialize("info")
	if err != nil {
		panic(err)
	}
	defer func(Log *zap.SugaredLogger) {
		err := Log.Sync()
		if err != nil {
			log.Errorf("ERROR: Error the logger syn: %v", err)
		}
	}(logger.Log)

	bridgeContractsProvider := app.New()

	err = bridgeContractsProvider.Listen(":8000")
	if err != nil {
		logger.Error("ContractsProvider::Main:", err)
		panic(err)
	}
}
