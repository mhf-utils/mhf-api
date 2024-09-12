package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"

	"mhf-api/config"
	"mhf-api/server/middlewares"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func Init(log *logger.Logger, newRelicApp *newrelic.Application) {
	log.Info("MHF-API:server:Init")

	router := mux.NewRouter()
	locales := []string{}

	localesConfig := map[string]string{
		"/en": config.GlobalConfig.Mhfdat.En.FilePath,
		"/fr": config.GlobalConfig.Mhfdat.Fr.FilePath,
		"/jp": config.GlobalConfig.Mhfdat.Jp.FilePath,
	}

	for path, filePath := range localesConfig {
		if len(filePath) > 0 {
			binaryFile := binary.GetBinaryFile(filePath)
			defer binaryFile.Close()

			subRouter := router.PathPrefix(path).Subrouter()
			middlewares.GetRouter(subRouter, log, binaryFile)
			locales = append(locales, path)
		}
	}

	routerKeeper := middlewares.RouterKeeper(log, locales)
	logging := middlewares.Logging(log, newRelicApp)

	log.Info("MHF-API listening on -> ", config.GlobalConfig.Host, config.GlobalConfig.Port)

	err := http.ListenAndServe(config.GlobalConfig.Port, middlewares.Chain(router, routerKeeper, logging))
	if err != nil {
		log.Fatal("An error occurred on starting the MHF-API")
	}
}
