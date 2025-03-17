package main

import (
	"bridge-eth-deployer/src/app"
	"bridge-eth-deployer/src/config"
	"bridge-eth-deployer/src/helpers/logger"
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

	bridgeEthDeployer := app.New()

	err = bridgeEthDeployer.Listen(":8000")
	if err != nil {
		log.Errorf("ERROR: Error starting the API server: %v", err)
		panic(err)
	}
}
