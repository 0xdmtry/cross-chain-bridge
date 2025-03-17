package main

import (
	"bridge-eth-compiler/src/app"
	"bridge-eth-compiler/src/config"
	"bridge-eth-compiler/src/helpers/logger"
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

	bridgeEthCompiler := app.New()

	err = bridgeEthCompiler.Listen(":8000")
	if err != nil {
		logger.Error("EthCompiler::Main:", err)
		panic(err)
	}
}
