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

	locales_config := map[string]config.MhfdatInfo{
		"/en": config.GlobalConfig.Mhfdat.En,
		"/fr": config.GlobalConfig.Mhfdat.Fr,
		"/jp": config.GlobalConfig.Mhfdat.Jp,
	}

	for locale, mhfdat_config := range locales_config {
		if mhfdat_config.Enable {
			binaryFile := binary.GetBinaryFile(mhfdat_config.FilePath)
			defer binaryFile.Close()

			subRouter := router.PathPrefix(locale).Subrouter()
			middlewares.GetRouter(subRouter, locale, log, binaryFile)
			locales = append(locales, locale)
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
