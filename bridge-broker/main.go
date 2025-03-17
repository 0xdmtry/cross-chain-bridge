package main

import (
	"bridge-broker/src/app"
	"bridge-broker/src/config"
	"bridge-broker/src/helpers/logger"
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
			log.Errorf("Broker::Main::Zap: %v", err)
		}
	}(logger.Log)

	bridgeBroker := app.New()

	err = bridgeBroker.Listen(":8000")
	if err != nil {
		logger.Error("Broker::Main:", err)
		panic(err)
	}
}
