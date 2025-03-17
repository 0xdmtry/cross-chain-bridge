package main

import (
	"bridge-accounts-creator/src/app"
	"bridge-accounts-creator/src/config"
	"bridge-accounts-creator/src/helpers/logger"
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

	bridgeEOA := app.New()

	err = bridgeEOA.Listen(":8000")
	if err != nil {
		log.Errorf("ERROR: Error starting the API server: %v", err)
		panic(err)
	}
}
