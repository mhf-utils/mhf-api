package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"

	"mhf-api/config"
	"mhf-api/server/middlewares"
	"mhf-api/utils/logger"
)

func Init(log *logger.Logger, newRelicApp *newrelic.Application) {
	log.Info("MHF-API:server:Init")

	router := mux.NewRouter()
	var prefixes []string

	router, prefixes = middlewares.GetLauncherRouter(log, router, prefixes)
	router, prefixes = middlewares.GetMhfdatRouter(log, router, prefixes)

	defer middlewares.CloseMhfdatBinaries()

	router_keeper := middlewares.RouterKeeper(log, prefixes)
	logging := middlewares.Logging(log, newRelicApp)

	log.Info("MHF-API listening on -> ", config.GlobalConfig.Host, config.GlobalConfig.Port)

	err := http.ListenAndServe(config.GlobalConfig.Port, middlewares.Chain(router, router_keeper, logging))
	if err != nil {
		log.Fatal("An error occurred on starting the MHF-API")
	}
}
