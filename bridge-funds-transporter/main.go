package main

import (
	"bridge-funds-transporter/src/app"
	"bridge-funds-transporter/src/config"
	"bridge-funds-transporter/src/helpers/logger"
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

	bridgeFundsTransporter := app.New()

	err = bridgeFundsTransporter.Listen(":8000")
	if err != nil {
		logger.Error("EthCompiler::Main:", err)
		panic(err)
	}
}
