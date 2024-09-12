package server

import (
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"

	"mhf-api/config"
	"mhf-api/server/middlewares"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func Init(log *logger.Logger, newRelicApp *newrelic.Application) {
	log.Info("MHF-API:server:Init")

	binary_file := binary.GetBinaryFile(config.GlobalConfig.Mhfdat.FilePath)
	defer binary_file.Close()

	router := middlewares.GetRouter(log, binary_file)
	logging := middlewares.Logging(log, newRelicApp)

	log.Info("MHF-API listening on -> ", config.GlobalConfig.Host, config.GlobalConfig.Port)

	err := http.ListenAndServe(config.GlobalConfig.Port, middlewares.Chain(router, logging))
	if err != nil {
		log.Fatal("An error occurred on starting the MHF-API")
	}
}
